"""
Sphinx extension to provide link anchors for configuration options.

Modeled after the standard :directive:`cmdoption` directive.
"""

from typing import Tuple, Dict, Iterator, Set, NamedTuple

from docutils.nodes import Element
from docutils.parsers.rst import directives
from sphinx import addnodes
from sphinx.addnodes import pending_xref
from sphinx.builders import Builder
from sphinx.directives import ObjectDescription
from sphinx.domains import Domain, ObjType
from sphinx.environment import BuildEnvironment
from sphinx.roles import XRefRole
from sphinx.util.nodes import make_refnode
from sphinx.util import logging


__licence__ = 'BSD (3 clause)'


logger = logging.getLogger(__name__)


ConfigOptionEntry = NamedTuple('ConfigOptionEntry',
                               [('name', str),
                                ('dispname', str),
                                ('typ', str),
                                ('docname', str),
                                ('anchor', str),
                                ('priority', str)])


class ConfigOptionRole(XRefRole):
    """Name of a configuration option, with automatic cross-reference."""

    def process_link(self, env: BuildEnvironment, refnode: Element,
                     has_explicit_title: bool, title: str, target: str) -> Tuple[str, str]:
        if not has_explicit_title:
            target = target.lstrip('~')  # only has a meaning for the title
            # if the first character is a tilde, don't display the section part of the contents
            if title[0:1] == '~':
                title = title[1:]
                dot = title.rfind('.')
                if dot != -1:
                    title = title[dot + 1:]
        return title, target


def alias_list(argument):
    """
    Convert whitespace-separated alias sigs into a list of strings.

    (Directive option conversion function.)
    """
    return argument.split()


class ConfigOptionDirective(ObjectDescription):
    """Name of a configuration option, usable as an external link target."""

    has_content = True
    required_arguments = 1
    option_spec = {
        'mandatory': directives.unchanged,
        'aliases': alias_list,
    }

    def handle_signature(self, sig, signode) -> Tuple[str, str]:
        parts = sig.split(sep='.', maxsplit=1)
        if len(parts) > 1:
            section, option = parts
        else:
            section, option = '', sig
        signode += addnodes.desc_name(text=option)
        if 'mandatory' in self.options:
            annotation = self.options['mandatory']
            if annotation:
                annotation = ' (mandatory: {})'.format(annotation)
            else:
                annotation = ' (mandatory)'
            signode += addnodes.desc_annotation(annotation, annotation)
        return section, option

    def add_target_and_index(self, name, sig, signode):
        anchor = 'config-option-%s' % sig.lower()
        signode['ids'].append(anchor)
        aliases = ['config-option-%s' % alias.lower()
                   for alias in self.options.get('aliases', [])]
        for alias_anchor in aliases:
            signode['ids'].append(alias_anchor)
        config = self.env.get_domain('stconf')
        config.add_config_option(sig, *name, anchor, location=signode)


class SyncthingConfigDomain(Domain):
    """Custom domain to group information regarding Syncthing's configuration."""

    name = 'stconf'
    label = 'Syncthing Configuration'
    directives = {
        'option': ConfigOptionDirective,
    }
    roles = {
        'opt': ConfigOptionRole(),
    }
    object_types = {
        'option': ObjType('option', 'opt'),
    }
    initial_data = {
        'sections': set(),  # string list
        'options': {},  # fullname -> docname, objtype
    }

    @property
    def config_sections(self) -> Set[str]:
        return self.data.setdefault('sections', [])

    @property
    def config_options(self) -> Dict[str, Tuple]:
        return self.data.setdefault('options', {})  # fullname -> (docname, node_id)

    def get_full_qualified_name(self, node):  # FIXME: what is this for?!
        return '{}.{}'.format('stconf-opt', node.arguments[0])

    def get_objects(self) -> Iterator[Tuple[str, str, str, str, str, int]]:
        for obj in self.config_options.values():
            yield obj

    def resolve_xref(self, env: BuildEnvironment, fromdocname: str, builder: Builder,
                     typ: str, target: str, node: pending_xref, contnode: Element
                     ) -> Element:
        searches = [target]
        if '.' not in target:
            searches += ['{}.{}'.format(section, target)
                         for section in self.config_sections]
        match = [(docname, anchor)
                 for name, sig, typ, docname, anchor, prio
                 in self.get_objects() if sig in searches]
        match = list(match)

        if len(match) > 0:
            todocname = match[0][0]
            targ = match[0][1]

            return make_refnode(builder, fromdocname, todocname, targ,
                                contnode, targ)
        logger.warning('Undefined config option: %s', target, location=node)
        return None

    def add_config_option(self, signature, section, option, anchor, location=None):
        """Add a new option anchor to the domain."""
        name = '{}.{}'.format('stconf-opt', signature)
        if section:
            self.config_sections.add(section)
        if signature in self.config_options:
            other = self.config_options[signature]
            logger.warning('Duplicate object description of %s, '
                           'other instance in %s, use :noindex: for one of them',
                           name, other.docname, location=location)
        self.config_options[signature] = ConfigOptionEntry(
            name=name,
            dispname=signature,
            typ='option',
            docname=self.env.docname,
            anchor=anchor, priority=0)


def setup(app):
    """Install the plugin.

    :param app: Sphinx application context.
    """
    app.add_domain(SyncthingConfigDomain)
    return
