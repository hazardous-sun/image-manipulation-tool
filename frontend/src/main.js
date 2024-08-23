import './style.css';
import './app.css';

import {EventsOn} from "../wailsjs/runtime";

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

