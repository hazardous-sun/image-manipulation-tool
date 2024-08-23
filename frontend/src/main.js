import './style.css';
import './app.css';

import logo from './assets/images/logo-universal.png';
import {Greet} from '../wailsjs/go/main/App';
import {EventsOn} from "../wailsjs/runtime";

// document.querySelector('#imagesPanel').innerHTML = `
// <!--    <img src="src/assets/images/cat3.jpg" alt="Image 1">-->
// <!--    <img src="src/assets/images/cat3.jpg" alt="Image 2">-->
// `;
// document.getElementById('logo').src = logo;

function setImage(path) {
    var originalImage = document.getElementById("originalImage");
    var previewImage = document.getElementById("previewImage");

    originalImage.src = path;
    previewImage.src = path;
}

EventsOn('set-image', (data) => {
    console.log(data.image);
    setImage("src/assets/temp/" + data.image);
});

