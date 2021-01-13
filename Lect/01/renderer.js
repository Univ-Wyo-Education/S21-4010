// Copyright (C) Philip Schlump, 2017.

window.$ = window.jQuery = require('./js/jquery-2.1.4.min.js');

// var theImage = document.getElementById("theImage");
var curImage = {};
var curHtml = {};
var isHtml = false;
var screen_width = 0;
var screen_height = 0;
var screen_ratio = 1.0;
var display_title = false;
var enlarge_image_to_fit = true;		// 'E' toggle 


//	Ratio of Image iX/y
//	Ratio of Screen X/y
//	if iX < sX && iY < sY - center image
//	if iX > sX && iY < sY - pad 1
//	if iX < sX && iY > sY - pad 2
//	if iX/y < sX/y
//	if iX/y > sX/y

theImage.onload = function() {
	let i_ratio = curImage.height / curImage.width;
	// console.log ( "Loaded image h/w/(h/w)", curImage.height, theImage.width, i_ratio, "screen w/h/(h/w)", screen_height, screen_width, screen_ratio );
	theImage.style.maxWidth = screen_width + "px";
	theImage.style.maxHeight = screen_height + "px";
	theImage.style.marginTop =  "0px";
	theImage.style.marginLeft =  "0px";
	theImage.style.marginBottom =  "0px";
	theImage.style.marginRight =  "0px";
	theImage.style.height = "auto";
	theImage.style.width = "auto";
	if ( curImage.height > screen_height && curImage.width > screen_width ) {
		if ( i_ratio > screen_ratio ) {
			// theImage.style.width = Math.floor(i_ratio * curImage.width) + "px";
			theImage.style.height = screen_height + "px";
			let x = Math.floor( screen_width - theImage.width ) / 2;
			theImage.style.marginLeft = x + "px";
			// console.log ( "i_ratio > screen_ratio, set height", screen_height, "marginLeft", x, theImage );
		} else {
			theImage.style.width = screen_width + "px";
			// theImage.style.height = Math.floor(i_ratio * curImage.width) + "px";
			let x = Math.floor( screen_height - theImage.height ) / 2;
			theImage.style.marginTop = x + "px";
			// console.log ( "i_ratio <= screen_ratio, set width", screen_width, "marginTop", x, theImage );
		}
	} else if ( curImage.height < screen_height && curImage.width > screen_width ) { // set max width - fit to secreen
		// console.log ( "set max width" );
		theImage.style.width = screen_width + "px";
		let x = Math.floor( screen_height - theImage.height ) / 2;
		theImage.style.marginTop = x + "px";
	} else if ( curImage.height > screen_height && curImage.width < screen_width ) { // set max height - fit to secreen
		// console.log ( "set max height" );
		theImage.style.height = screen_height + "px";
		let x = Math.floor( screen_width - theImage.width ) / 2;
		theImage.style.marginLeft = x + "px";
	} else {
		if ( enlarge_image_to_fit ) {
			if ( i_ratio > screen_ratio ) {
				// theImage.style.width = Math.floor(i_ratio * curImage.width) + "px";
				theImage.style.height = screen_height + "px";
				let x = Math.floor( screen_width - theImage.width ) / 2;
				theImage.style.marginLeft = x + "px";
				// console.log ( "ENLARGE: i_ratio > screen_ratio, set height", screen_height, "marginLeft", x, theImage );
			} else {
				theImage.style.width = screen_width + "px";
				// theImage.style.height = Math.floor(i_ratio * curImage.width) + "px";
				let x = Math.floor( screen_height - theImage.height ) / 2;
				theImage.style.marginTop = x + "px";
				// console.log ( "ENLARGE: i_ratio <= screen_ratio, set width", screen_width, "marginTop", x, theImage );
			}
		} else {
			// padding needed - center image. - or expand image to fit screen?
			if ( curImage.width < (screen_width-2) ) {
				let x = Math.floor( screen_width - theImage.width ) / 2;
				theImage.style.marginLeft = x + "px";
			}
			if ( curImage.height < (screen_height-2) ) {
				let x = Math.floor( screen_height - theImage.height ) / 2;
				theImage.style.marginTop = x + "px";
			}
		}
	}
}

// console.log ( "document.getElementById()=", theImage.src );
// x.src = imgArr[imgNo].image;

var {ipcRenderer, remote} = require('electron');  
var main = remote.require("./main.js");

// Listen for async-reply message from main process
ipcRenderer.on('show-img', (event, arg) => {  
	let $ = window.jQuery;

	console.log('Got show-img', arg);

	$("#theImageDiv").show();
	$("#theHtmlDiv").hide();

	theImage.src = arg.image;
	curImage = arg;
	isHtml = false;

	showTitle();

});

ipcRenderer.on('show-html', (event, arg) => {  
	let $ = window.jQuery;

	console.log('Got show-html', arg);

	$("#theImageDiv").hide();
	$("#theHtmlDiv").show();

	// theImage.src = arg.image;
	$("#theHtml").html(arg.body);
	curImage = {};
	curHtml = arg;
	isHtml = true;

	showTitle();

});

function showTitle() {  
	if ( display_title ) {
		if ( isHtml ) {
			$("#theTitle").html("<h1 class=\"html\">"+curHtml.title+"</h1>");
		} else {
			$("#theTitle").html("<h1 class=\"image\">"+curImage.title+"</h1>");
		}
		$("#theTitle").show();
	} else {
		$("#theTitle").hide();
	}
}

function showHelp() {  
	let $ = window.jQuery;
	let helpMsg = `
<div class="slide">

<h1> Help - Commands </h1>

<table class="table table-bordered">
	<thead>
		<tr> <th> Key </th> <th> Description </th> </tr>
	</thead>
	<tbody>
		<tr> <td>  space </td>        <td> Move to next slide </td> </tr>
		<tr> <td>  right-arrow </td>  <td> Move to next slide </td> </tr>
		<tr> <td>  left-arrow </td>   <td> Move to previous slide </td> </tr>
		<tr> <td>  t </td>            <td> Toggle title display</td> </tr>
	</tbody>
</table>

</div>
`;

	console.log('Got showHelp');

	$("#theImageDiv").hide();
	$("#theHtmlDiv").show();

	// theImage.src = arg.image;
	$("#theHtml").html(helpMsg);
	curImage = {};
	curHtml = { "html": "*help-msg*", "body": helpMsg, "title": "Help Info" };
	isHtml = true;

	showTitle();
}

ipcRenderer.on('set-screen', (event, arg) => {  
	// console.log('Got set-screen', arg);
	screen_height = 0+arg.height;
	screen_width = 0+arg.width;
	screen_ratio = screen_height/screen_width;
});

$(document).on('keydown', function(e) {
	// console.log("keydown", e.which);
	switch ( e.which ) {
	case 39: // Next: 39
	case 32: // Next: Space
		ipcRenderer.send('get-next-img', 0);	
		break;
	case 37: // Prev: 37
		ipcRenderer.send('get-prev-img', 0);	
		break;
	case 84: // t: 37
		display_title = ! display_title ;
		// console.log ( "display_title=", display_title );
		showTitle();
		break;
	case 69: // E: 69
		if ( e.shiftKey ) {
			enlarge_image_to_fit = ! enlarge_image_to_fit ;
		}
		break;
	case 88: // alt-X: - Open the dev tools
		if ( e.altKey ) {
			ipcRenderer.send( 'open-dev-tools' , 0);	
		}
		break;
	case 191: // ? - show help slide
		if ( e.shiftKey ) {
			showHelp();
		}
		break;
	// 'Q' - exit
	// r - rotate 90
	// 'D' - delete file - confirm dialog
	default:
		console.log("keydown", e.which, e);
	}
});

// Send async message to main process
ipcRenderer.send('get-0th-img', 0);			

