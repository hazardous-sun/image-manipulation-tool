import './style.css';
import './app.css';

import {EventsEmit, EventsOn} from "../wailsjs/runtime";
import {GrayScale, Transform} from "../wailsjs/go/main/App";

// Theme ---------------------------------------------------------------------------------------------------------------

EventsOn('set-theme', (data) => async function () {
    for (let i = 0; i < 1000; i++) {
        console.log("JS theme triggered")
    }
    let themeName = getThemeData()
});

async function getThemeData() {
    try {
        const response = await fetch('config.json');
        const data = await response.json();

        const themePath = `src/assets/themes/${data.theme}`;
        const themeResponse = await fetch(themePath);
        const themeData = await themeResponse.json();

        // Use themeData to apply the theme to your application
        applyTheme(themeData);
    } catch (error) {
        console.error('error when loading theme:', error);
    }
}

function applyTheme(themeData) {
    document.getElementById("app")
        .style
        .backgroundColor = themeData["sidebar-color"]
    document.getElementById("imagesPanel")
        .style
        .backgroundColor = themeData["background-color"]
    document.querySelectorAll('.sideBar, li')
        .forEach(element => {
            element.style.borderColor = themeData["sidebar-border-color"]
        });
    document.querySelectorAll('.sideBar, ul, li, a')
        .forEach(element => {
            element.style.color = themeData["font-color"]
        })
    document.querySelectorAll('.sideBar, ul, li, ul.dropdown')
        .forEach(element => {
            element.style.backgroundColor = themeData["dropdown-menu-color"]
        })
}

// Image handling ------------------------------------------------------------------------------------------------------

EventsOn('set-origin-prev', (data) => {
    setOriginPrev(data.path);
});

EventsOn('set-prev', (data) => {
    setPrev(data.path);
});

function setOriginPrev(path) {
    let originalImage = document.getElementById("originalImage");
    let previewImage = document.getElementById("previewImage");
    originalImage.src = "";
    previewImage.src = "";
    originalImage.src = "src/assets/temp/origin/" + path;
    previewImage.src = "src/assets/temp/prev/" + path;
}

function setPrev(path) {
    let previewImage = document.getElementById("previewImage");
    previewImage.src = "";
    previewImage.src = "src/assets/temp/prev/" + path;
}

EventsOn('get-prev', () => {
    let previewImage = document.getElementById("previewImage").src.toString();
    EventsEmit("receive-prev", {"path": previewImage})
})

// Geometric transformations -------------------------------------------------------------------------------------------

function geoTransform(code, x, y) {

    let prevImageSrc = document.getElementById('previewImage').src

    if (prevImageSrc === "") {
        return
    }

    Transform(prevImageSrc, code, x, y) 
}

function getXY(baseValue) {
    let x = document.getElementById('xAxis').value

    if (x === "") {
        x = baseValue
    } else {
        x = Number(x)
    }

    let y = document.getElementById('yAxis').value

    if (y === "") {
        y = baseValue
    } else {
        y = Number(y)
    }

    return [x, y]
}

window.imgTranslate = function () {
    let values = getXY(0)
    geoTransform(0, values[0], values[1])
}

window.imgResize = function () {
    let values = getXY(1)
    geoTransform(1, values[0], values[1])
}

window.mirrorH = function () {
    geoTransform(2, 0, 0)
}

window.mirrorV = function () {
    geoTransform(3, 0, 0)
}

window.imgRotate = function () {
    let values = getXY(0)
    geoTransform(4, values[0], values[1])
}

// Filter --------------------------------------------------------------------------------------------------------------

// Grayscale
window.filterGrayScale = function () {
    // send from [29::]
    let prevImageSrc = document.getElementById('previewImage').src
    GrayScale(prevImageSrc).then()
}
