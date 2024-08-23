import './style.css';
import './app.css';

import {EventsOn} from "../wailsjs/runtime";

EventsOn('set-theme', (data) => async function () {
    let themeName = getThemeData()
});

async function getThemeData() {
    try {
        const response = await fetch('your_initial_json_file.json');
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

EventsOn('set-image', (data) => {
    setImage(data.fileExt);
});

function setImage(fileExt) {
    var originalImage = document.getElementById("originalImage");
    var previewImage = document.getElementById("previewImage");

    originalImage.src = "src/assets/temp/origin" + fileExt;
    previewImage.src = "src/assets/temp/prev" + fileExt;
}

