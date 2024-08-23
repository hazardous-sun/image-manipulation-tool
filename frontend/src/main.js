import './style.css';
import './app.css';

import {EventsOn} from "../wailsjs/runtime";

function setImage(fileExt) {
    var originalImage = document.getElementById("originalImage");
    var previewImage = document.getElementById("previewImage");

    originalImage.src = "frontend/src/assets/temp/origin" + fileExt;
    previewImage.src = "frontend/src/assets/temp/origin" + fileExt;
}

EventsOn('set-image', (data) => {
    console.log(data.fileExt);
    setImage(data.fileExt);
});

