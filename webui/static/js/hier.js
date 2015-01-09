if (typeof Hier === "undefined") Hier = {};

Hier.sync_resolutions = function(status, resolution) {
    var hider = null;
    for (hider = resolution; hider !== null; hider = hider.parentElement) {
	if (hider.classList.contains("hider")) {
	    break;
	}
    }
    if (hider == null)
	return;
    var substatuses = Hier.statuses[status.value];
    if (status.value === "") {
	hider.hidden = true;
    } else if (substatuses) {
	hider.hidden = false;
	resolution.innerHTML = "";
	substatuses.forEach(function(st) {
	    var el = document.createElement("option");
	    el.setAttribute("value", st);
	    el.innerText = st;
	    resolution.appendChild(el);
	});
    } else {
	hider.hidden = true;
    }
}
