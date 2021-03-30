let reloaded = 0;

let hotReloader = function() {
let socket = new WebSocket("ws://127.0.0.1:9023/ws");
console.log("Attempting Connection...");

socket.onopen = () => {
    console.log("Successfully Connected");
};

socket.onmessage = (m) => {
    reload()

   console.log("Changes detected ", m);
};

socket.onclose = event => {
    console.log("Socket Closed Connection: ", event);
};

socket.onerror = error => {
    console.log("Socket Error: ", error);

};
}

function debounce(func, wait, immediate) {
	var timeout;
	return function() {
		var context = this, args = arguments;
		var later = function() {
			timeout = null;
			if (!immediate) func.apply(context, args);
		};
		var callNow = immediate && !timeout;
		clearTimeout(timeout);
		timeout = setTimeout(later, wait);
		if (callNow) func.apply(context, args);
	};
};

let reload = debounce(function () { location.reload() }, 1500)

hotReloader()