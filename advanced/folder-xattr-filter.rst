xattrFilter
===========

The extended attributes filter is used to manage which attributes should
be synced and which ones should not be.

When the filter shows
---------------------

The ability to manipulate the extended attributes filter in the advanced
panel of a folder configuration only appears after enabling the 
``Sync Extended Attributes`` option.

Default Behaviour
-----------------

When setting filters it's good to understand the default behaviour, this
default behaviour changes once you add rules. When no rules are present
everything is permitted. However, once a single rule is set this default
behaviour flips to denying everything that does not match a rule.

These rules work on a 'first-match'-principle. This means that once it
matches any rule (going from top to bottom), that rule is applied and the
other rules won't be checked further. When no rules result in a match,
then the default is deny unless stated otherwise (or when no rules are
present at all. In that case the default is to permit everything).

To override this default behaviour when adding rules, you should add a
rule in the bottom which permits everything (*).

The GUI reminds you of the default behaviour when you don't have an any-rule
set.

Setting rules
-------------

To set a rule press the button 'Add new rule' to append an empty rule to the
list. The check-box indicates whether to permit (checked) or deny (unchecked)
the matching pattern and the text-field is the pattern to be matched. Adding
a rule while already having a wild-card (*) added will result in the new rule
being added as second-last item. When saving the configuration, empty rules
will be removed. Closing the window without pressing save will cause your
made changes to be reset.

Be aware that the rules aren't being validated on correctness beyond the point
of cleaning out empty rules.

Removing rules
--------------

To remove a rule simply press on the trash bin on the right hand side of the
rule that you want to remove.

.. seealso:: :doc:`folder-send-xattrs`, :doc:`folder-sync-xattrs`
