POST /rest/system/resume
========================

Resume the given device or all devices.

Takes the optional parameter ``device`` (device ID). When omitted,
resumes all devices.  Returns status 200 and no content upon success, or status
500 and a plain text error on failure.
