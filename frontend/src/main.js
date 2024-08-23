import './style.css';
import './app.css';

import {EventsOn} from "../wailsjs/runtime";

function setImage(fileExt) {
    var originalImage = document.getElementById("originalImage");
    var previewImage = document.getElementById("previewImage");

    originalImage.src = "src/assets/temp/origin" + fileExt;
    previewImage.src = "src/assets/temp/prev" + fileExt;
}

EventsOn('set-image', (data) => {
    console.log(data.fileExt);
    setImage(data.fileExt);
});

