var collator = new Intl.Collator(undefined, {numeric: true, sensitivity: 'base'});

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

function stripVersionPath(path, versions) {
    var slash = path.indexOf('/', 1);
    if (slash != -1) {
        if (versions.indexOf(path.slice(1, slash)) != -1) {
            path = path.slice(slash);
        }
    }
    return path;
}

function redirectToPath(newPath) {
    const fragment = window.location.href.indexOf('#');
    if (fragment != -1) {
        newPath += window.location.href.slice(fragment);
    }

    if (newPath && newPath != window.location.pathname) {
        window.location.replace(newPath);
    }
}

function redirectToVersion(target, available) {
    const tail = stripVersionPath(window.location.pathname, available + [target]);

    var newPath = '';
    if (target) {
        newPath += '/' + target;
    }
    if (tail) {
        newPath += tail;
    }
    redirectToPath(newPath);
}


const urlParams = new URLSearchParams(window.location.search);
const versionParam = urlParams.get('version');


if (versionParam) {
    getVersions.then(function (available) {
        const useVersion = findBestVersion(versionParam, available);
        redirectToVersion(useVersion, available);
    });
}
