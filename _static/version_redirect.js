var collator = new Intl.Collator(undefined, {numeric: true, sensitivity: 'base'});

// Override MIME type to avoid "invalid XML" warnings when using $.getJSON()
$.ajaxSetup({beforeSend: function (xhr) {
    if (xhr.overrideMimeType) {
        xhr.overrideMimeType("application/json");
    }
}});

const VERSIONS_LIST = "/versions.json";

const getVersions = $.getJSON(VERSIONS_LIST).then(function (data) {
    // Start with highest version number, using natural sorting
    data.entries.sort(collator.compare).reverse();
    return data.entries;
});

function findBestVersion(version, available) {
    var bestVersion = '';
    available.some(function (candidate) {
        if (version.startsWith(candidate)) {
            // Direct prefix match
            bestVersion = candidate;
            return true;
        }
        if (collator.compare(candidate, version) < 0) {
            // Available version is numerically lower than requested
            if (version.startsWith(candidate.slice(0, candidate.lastIndexOf('.')))) {
                // Use the lower version if it only differs in last component
                bestVersion = candidate;
            }
            // Stop checking even older versions
            return true;
        }
        bestVersion = candidate;
        return false;
    });
    // Filter out any higher versions which differ in more than the last component
    if (!version.startsWith(bestVersion.slice(0, bestVersion.lastIndexOf('.')))) {
        bestVersion = '';
    }
    return bestVersion;
}

function splitVersionPath(path, versions) {
    // Find end of first path component, disregarding leading slash
    var slash = path.indexOf('/', 1);
    if (slash != -1) {
        var firstComponent = path.slice(1, slash);
        if (versions.indexOf(firstComponent) != -1) {
            // Component is a valid known version path, split it off
            return [firstComponent, path.slice(slash)];
        }
    }
    return ['', path];
}

function redirectToPath(newPath, keepHistory) {
    const fragment = window.location.href.indexOf('#');
    if (fragment != -1) {
        newPath += window.location.href.slice(fragment);
    }

    if (newPath && newPath != window.location.pathname) {
        if (keepHistory) {
            window.location.assign(newPath);
        } else {
            window.location.replace(newPath);
        }
    }
}

function redirectToVersion(target, available, keepHistory) {
    const tail = splitVersionPath(window.location.pathname, available + [target])[1];

    var newPath = '';
    if (target) {
        newPath += '/' + target;
    }
    if (tail) {
        newPath += tail;
    }
    redirectToPath(newPath, keepHistory);
}

function setVersionPickerOptions() {
    getVersions.then(function (available) {
        var items = [
            '<option value="">latest</option>'
        ];
        var current = splitVersionPath(window.location.pathname, available)[0];
        $.each(available, function (key, val) {
            var item = '<option value="' + val + '"';
            if (val == current) item += ' selected';
            item += '>' + val + '</option>';
            items.push(item);
        });
        let sel = document.createElement('select');
        sel.setAttribute('id', 'version-picker');
        sel.setAttribute('style', 'font: inherit;');
        sel.onchange = pickVersion;
        sel.innerHTML = items.join('');
        let note = document.createElement('div');
        note.classList.add('admonition', 'hint');
        note.append('Browsing documentation for version: ');
        note.append(sel);
        var doc = document.getElementsByClassName('document')[0];
        doc.prepend(note);
    });
}

function pickVersion() {
    getVersions.then(function (available) {
        const targetVersion = document.getElementById('version-picker').value;
        redirectToVersion(targetVersion, available, true);
    });
}


const urlParams = new URLSearchParams(window.location.search);
const versionParam = urlParams.get('version');


if (versionParam) {
    getVersions.then(function (available) {
        const useVersion = findBestVersion(versionParam, available);
        redirectToVersion(useVersion, available, false);
    });
}

window.addEventListener('DOMContentLoaded', setVersionPickerOptions);
