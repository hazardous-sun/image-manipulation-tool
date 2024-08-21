import './style.css';
import './app.css';

import logo from './assets/images/logo-universal.png';
import {Greet} from '../wailsjs/go/main/App';

document.querySelector('#imagesPanel').innerHTML = `
<!--    <img id="logo" class="logo">-->
<!--      <div class="result" id="result">Please enter your name below 👇</div>-->
<!--      <div class="input-box" id="input">-->
<!--        <input class="input" id="name" type="text" autocomplete="off" />-->
<!--        <button class="btn" onclick="greet()">Greet</button>-->
<!--      </div>-->
<!--    </div>-->

<div class="image-container centerHV">
    <img src="src/assets/images/cat3.jpg" alt="Image 1">
    <img src="src/assets/images/cat3.jpg" alt="Image 2">
</div>
`;
document.getElementById('logo').src = logo;

let nameElement = document.getElementById("name");
nameElement.focus();
let resultElement = document.getElementById("result");

// Setup the greet function
window.greet = function () {
    // Get name
    let name = nameElement.value;

    let tempValue = ""

    // Check if the input is empty
    if (name === "") return;

    // Call App.Greet(name)
    try {
        Greet(name)
            .then((result) => {
                // Update result with data back from App.Greet()
                tempValue += result;
            })
            .catch((err) => {
                console.error(err);
            });
    } catch (err) {
        console.error(err);
    }
};
