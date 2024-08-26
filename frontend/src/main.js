import './style.css';
import './app.css';

import {EventsOn} from "../wailsjs/runtime";
import {GrayScale} from "../wailsjs/go/main/App";

// Theme ---------------------------------------------------------------------------------------------------------------

EventsOn('set-theme', (data) => async function () {
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
        console.error('Error loading theme:', error);
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

// Filter --------------------------------------------------------------------------------------------------------------

// Grayscale
window.filterGrayScale = function () {
    // send from [29::]
    let prevImageSrc = document.getElementById('previewImage').src
    GrayScale(prevImageSrc).then()
}
