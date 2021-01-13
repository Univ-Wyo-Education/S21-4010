// Copyright (C) Philip Schlump, 2017.

const electron = require('electron');
// Module to control application life.
const app = electron.app
const {protocol, ipcMain} = require('electron')
const {globalShortcut} = require('electron');

// Module to create native browser window.
const BrowserWindow = electron.BrowserWindow

const path = require('path');
const url = require('url');
const fs = require('fs');
// const readJsonSync = require('read-json-sync');


const sizeOf = require('image-size');



var itemNo = 0;
var itemInit = false;
var itemArr = {
	"items": [
	 	{ "html": "./slide/slide00.html"											, "title": "Why - Blockchain" }
	,	{ "html": "./slide/slide01.html"											, "title": "Example Ethereum/Blockchain Jobs" }
	,	{ "html": "./slide/slide01a.html"											, "title": "Changes to Go in last 5 years" }
	,	{ "html": "./slide/slide02.html"											, "title": "Languages Used in Jobs" }
	,	{ "html": "./slide/slide00a.html"											, "title": "Not Lost" }

	,	{ "image": "./img/dsc_025_0249.jpg"											, "title": "View out our front door" }
	,	{ "image": "./img/dsc-26950-cambridge-beach-04.jpg"							, "title": "Our home four 5 years" }
	,	{ "image": "./img/dsc_163_0680ed1_vg_at_anchor_eluthra.jpg"					, "title": "Palm trees and picnick bench" }
	]
};

//	,	{ "image": "./img/dscn-23094-6534-agii-at-anchor.jpg"						, "title": "Remote exotic locations" }
//	,	{ "image": "./img/pict_044_0025.jpg"										, "title": "Swim with fish" }
//	,	{ "image": "./img/img_144_3237ed1-dolphins-and-agii.jpg"					, "title": "And dolphins" }
//	,	{ "image": "./img/dsc-27436-055_0120-agii-at-anchor-treasure-rainbow.jpg"	, "title": "At the end of the rainbow" }

	// ,	{ "image": "./img/img_174_5721.jpg"											, "title": "Swim with fish" }
	//,	{ "image": "./img/spinner-gif-13.gif"										, "title": "spinner" }
//	,	{ "html": "./slide/slide01.html"											, "title": "YCombinator. Functional Jobs" }
//	,	{ "html": "./slide/slide02.html"											, "title": "YCombinator. React/Redux Jobs" }
//
//	,	{ "image": "./img/ReactReduxFlow.png"										, "title": "Redux Flow" }
//	,	{ "html": "./slide/slide04.html"											, "title": "An Ation Creator" }
//	,	{ "html": "./slide/slide05.html"											, "title": "A Reducer" }
//
//	,	{ "image": "./img/model.png"												, "title": "Model - Finance" }
//	,	{ "image": "./img/pict_044_0089.jpg"										, "title": "Swimming with sharks" }
//
//	,	{ "html": "./slide/slide99.html"											, "title": "By Philip Schlump" }

const imgFormats = [ ".bmp", ".cur", ".gif", ".ico", ".jpeg", ".jpg", ".png", ".psd", ".tiff", ".webp", ".svg", ".dds" ];
const htmlFormats = [ ".html" ];
const jsonFormats = [ ".json", ".img-vwr" ];
const xFormats = imgFormats.concat( htmlFormats );

// console.log('process.argv=', process.argv, process.argv.length );
if ( process.argv.length > 2 ) { // if len > 2 => have args/fns to process.
	// console.log ( "Has Args" );
	for ( let j = 2, mx = process.argv.length; j < mx; j++ ) {
		// console.log ( "Args["+j+"]="+process.argv[j] );
		// split file name into extension, name, - convert to hard path, pick off the dirname.
		let fn = process.argv[j];
		let ext = path.extname(fn).toLowerCase();				// index.html -> .html
		let dn = path.dirname(fn);				// ./index.html -> ., ./xyz/index.html -> ./xyz
		let bn = path.basename(fn);				// ./index.html -> index.html, ./xyz/index.html -> index.html
		// if is an image or htmlFormats
		// console.log ( "fn=", fn, " ext=", ext, " dn=", dn, " bn=", bn );
		if ( isDir ( fn ) ) {
			let fns = readDir ( fn );
			// console.log ( "fns from readDir() =", fns );
			// 	get list of files matching in direcotry - read directory - populate itemArr
			fns = matchExt ( fns, xFormats );
			// console.log ( "fns after matchExt() =", fns );
			let pos = findFn ( bn, fns ); //	find this file in directory - set itemNo
			// console.log ( "pos findFn() =", pos );
			itemNo = 0;
			if ( ! itemInit ) {
				itemArr = { items: [] };
				itemInit = true;
			}
			for ( let k = 0, my = fns.length; k < my; k++ ) {
				itemArr.items.push ( { "image": fns[k], "title": path.basename(fns[k]) } );		
			}
		}
		else if ( inArr ( ext, xFormats ) >= 0 ) {		// imgFormats, htmlFormats
			//	then - pick off "directory" 
			// console.log ( "inArr is true, ext matched" );
			let fns = readDir ( dn );
			// console.log ( "fns from readDir() =", fns );
			// 	get list of files matching in direcotry - read directory - populate itemArr
			fns = matchExt ( fns, xFormats );
			// console.log ( "fns after matchExt() =", fns );
			let pos = findFn ( bn, fns ); //	find this file in directory - set itemNo
			// console.log ( "pos findFn() =", pos );
			itemNo = pos;
			if ( ! itemInit ) {
				itemArr = { items: [] };
				itemInit = true;
			}
			for ( let k = 0, my = fns.length; k < my; k++ ) {
				itemArr.items.push ( { "image": fns[k], "title": path.basename(fns[k]) } );		
			}
			// console.log ( itemArr );
		}
		// if is a .json/img-vwr file
		// 	read in file into itemArr, set itemNo = 0
		else if ( inArr ( ext, jsonFormats ) >= 0 ) {		// imgFormats, htmlFormats
			// console.log ( "Read in JSON File ***********************************", fn );
			itemArr = readJsonSync(fn);
			// console.log ( itemArr );
		}
	}
	//	process.exit(0);		// xyzzy - testing
}

// isDir returns a synchronous True if fn is a directory
function isDir ( fn ) {
	try {
		stats = fs.lstatSync(fn); // Query the entry
		// Is it a directory?
		if (stats.isDirectory()) {
			return true; // Yes it is
		}
	} catch (e) {
		// return false; 	// fall through to false
	}
	return false; 
}

// readJsonSync reads and parses a JSON file synchronously.
function readJsonSync(filePath) {
	let s = String(fs.readFileSync(filePath, "utf8"));
	return JSON.parse(s);
};

// inArr returns -1 if "it" is not found in array "arr".  If found it returns the subscript.
function inArr ( it, arr ) {
	// console.log ( "it=["+it+"] arr="+JSON.stringify(arr));
	for ( let i = 0, mx = arr.length; i < mx; i++ ) {
		if ( arr[i] === it ) {
			return ( i );
		}
	}
	return -1;
}

// let fns = readDir ( dn );
// Note: https://code-maven.com/list-content-of-directory-with-nodejs
function readDir ( dirpath ) {
	var files = fs.readdirSync(dirpath).sort();		// returns sorted array of names.
	// process list to just names, no dirs
	for ( let i = 0, mx = files.length; i < mx; i++ ) { 
		files[i] = dirpath + path.sep + files[i];
	}
	return files;
}
			
// let pos = findFn ( bn, fns ); //	find this file in directory - set itemNo
function findFn ( name, arr ) {
	for ( let i = 0, mx = arr.length; i < mx; i++ ) {
		let bn = path.basename ( arr[i] );
		if ( bn === name ) {
			return i;
		}
	}
	return 0;
}
			
// fns = matchExt ( fns, imgFormats );
function matchExt ( arrOfFn, validExt ) {
	let rv = [];
	for ( let fni = 0, my = arrOfFn.length; fni < my; fni++ ) {
		let ext = path.extname(arrOfFn[fni]).toLowerCase();
		if ( inArr ( ext, validExt ) >= 0 ) {
			rv.push ( arrOfFn[fni] );
		}
	}
	return rv;
}

function getSizes ( itemArr ) {
	for ( let i = 0, mx = itemArr.items.length; i < mx; i++ ) {
		if ( itemArr.items[i].image ) {
			let fn = itemArr.items[i].image;
			try {
				let dimensions = sizeOf( fn );
				console.log(fn,dimensions.width, dimensions.height);
				itemArr.items[i].height = dimensions.height;
				itemArr.items[i].width = dimensions.width;
			}catch(e){
				console.log("Error with:", fn );
			}
		}
		if ( itemArr.items[i].html ) {
			let fn = itemArr.items[i].html;
			let s = String(fs.readFileSync(fn, "utf8"));
			console.log(fn,"HTML length=",s.length);
			itemArr.items[i].body = s;
		}
	}
}

getSizes ( itemArr ) ;

// Keep a global reference of the window object, if you don't, the window will
// be closed automatically when the JavaScript object is garbage collected.
let mainWindow;
let s_width, s_height;

function createWindow () {
	const {width, height} = electron.screen.getPrimaryDisplay().workAreaSize;
	s_width = width;
	s_height = height;

	// Create the browser window.
	mainWindow = new BrowserWindow({width: width, height: height, backgroundColor: '#000'});

	// and load the index.html of the app.
	mainWindow.loadURL(url.format({
		pathname: path.join(__dirname, 'index.html'),
		protocol: 'file:',
		slashes: true
	}))

	// Open the DevTools.
	// mainWindow.webContents.openDevTools();		// PJS - moded to include.

	// Emitted when the window is closed.
	mainWindow.on('closed', function () {
		// Dereference the window object, usually you would store windows
		// in an array if your app supports multi windows, this is the time
		// when you should delete the corresponding element.
		mainWindow = null
	});
}

app.on('ready', createWindow);

// ========================================================================================================================================================
// ========================================================================================================================================================
// ========================================================================================================================================================

// --------------- pjs --------------------------------------------------------------------------------------


function nextImage () {
	// console.log ( "itemNo", itemNo, "len()", itemArr.items.length ) ;
	if ( itemArr.items[itemNo].image ) {
		mainWindow.webContents.send('show-img', itemArr.items[itemNo]);
	}
	if ( itemArr.items[itemNo].html ) {
		mainWindow.webContents.send('show-html', itemArr.items[itemNo]);
	}
	itemNo++;
	if ( itemNo >= itemArr.items.length ) {
		itemNo = 0;
	}
}

function prevImage () {
	itemNo--;
	if ( itemNo < 0 ) {
		itemNo = itemArr.items.length-1;
	}
	// console.log ( "itemNo", itemNo, "len()", itemArr.items.length ) ;
	if ( itemArr.items[itemNo].image ) {
		mainWindow.webContents.send('show-img', itemArr.items[itemNo]);
	}
	if ( itemArr.items[itemNo].html ) {
		mainWindow.webContents.send('show-html', itemArr.items[itemNo]);
	}
}


app.on('ready', () => {
	protocol.registerFileProtocol('atom', (request, callback) => {
		const url = request.url.substr(7)
		callback({path: path.normalize(`${__dirname}/${url}`)})
	}, (error) => {
		if (error) {
			console.error('Failed to register protocol');
		}
	})
})

// -------------------------------------------------------------------------------------------------------------
// Listen for sync message from renderer process
ipcMain.on('get-0th-img', (event, arg) => {
    // console.log('received get-0th-img', arg);
	mainWindow.webContents.send('set-screen', { "height": s_height, "width": s_width } );
	nextImage();
//	if ( itemArr.items[itemNo].image ) {
//		mainWindow.webContents.send('show-img', itemArr.items[itemNo]);
//	}
//	if ( itemArr.items[itemNo].html ) {
//		mainWindow.webContents.send('show-html', itemArr.items[itemNo]);
//	}
//	itemNo++;
//	if ( itemNo >= itemArr.items.length ) {
//		itemNo = 0;
//	}
});

//		ipcRenderer.send('get-next-img', 0);
ipcMain.on('get-next-img', (event, arg) => {
    // console.log('received get-next-img', arg);
	nextImage();
});

//		ipcRenderer.send('get-prev-img', 0);
ipcMain.on('get-prev-img', (event, arg) => {
    // console.log('received get-prev-img', arg);
	prevImage();
});

ipcMain.on('open-dev-tools', (event, arg) => {
	mainWindow.webContents.openDevTools();		// PJS - moded to include.
});

// --------------- pjs --------------------------------------------------------------------------------------

// Quit when all windows are closed.
app.on('window-all-closed', function () {
	// On OS X it is common for applications and their menu bar to stay active until the user quits explicitly with Cmd + Q
	// if (process.platform !== 'darwin') { // PJS removed - if mac then just close like wincows.
		app.quit()
	// }
})

app.on('activate', function () {
	// On OS X it's common to re-create a window in the app when the
	// dock icon is clicked and there are no other windows open.
	if (mainWindow === null) {
		createWindow()
	}
})

