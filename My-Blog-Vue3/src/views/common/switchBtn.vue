<template>
  <div>
    <input id="switchBtnRef" type="checkbox" />
    <button>
      <div class="🌞">
        <div class="🌛">
          <div class="🌕"></div>
          <div class="🌕"></div>
          <div class="🌕"></div>
        </div>
      </div>
      <img class="✈️" src="../../assets/file/switch-feiji.png" />
      <img class="🐻" src="../../assets/file/switch-guaishou.png" />
      <div class="⭕️"></div>
      <div class="💭">
        <div class="☁️"></div>
        <div class="☁️"></div>
        <div class="☁️"></div>
        <div class="☁️"></div>
        <div class="☁️"></div>
        <div class="☁️"></div>
        <div class="☁️"></div>
        <div class="☁️"></div>
      </div>
      <div class="🌟">
        <div class="✨"></div>
        <div class="✨"></div>
        <div class="✨"></div>
        <div class="✨"></div>
        <div class="✨"></div>
        <div class="✨"></div>
        <div class="✨"></div>
        <div class="✨"></div>
        <div class="✨"></div>
        <div class="✨"></div>
        <div class="✨"></div>
      </div>
    </button>
  </div>
</template>

<script setup lang="ts"></script>

<style scoped>
input {
  position: absolute;
  top: 100px;
  left: 50%;
  width: 390px;
  height: 146px;
  z-index: 99;
  opacity: 0;
  transform: translateX(-50%);
}
button {
  margin-top: 100px;
  margin-left: 50%;
  width: 390px;
  height: 146px;
  transform: translateX(-50%);
  box-shadow: calc(var(--width) * 0) calc(var(--width) * 0.02) calc(var(--width) * 0.01)
      calc(var(--width) * -0.0025) hsl(210 10% 100% / 0.95),
    calc(var(--width) * 0) calc(var(--width) * -0.02) calc(var(--width) * 0.01)
      calc(var(--width) * -0.0025) hsl(210 10% 10% / 0.2),
    calc(var(--width) * 0) calc(var(--width) * 0.02) calc(var(--width) * 0.5) 0
      hsl(210 10% 100% / 0.15);
  border-radius: 146px;
  border: 0;
  padding: 12px;
  position: relative;
  overflow: hidden;
}
button::after {
  content: "";
  position: absolute;
  inset: 0;
  box-shadow: calc(var(--width) * 0) calc(var(--width) * -0.025) calc(var(--width) * 0.025) 0
      #17191c26 inset,
    calc(var(--width) * 0) calc(var(--width) * 0.025) calc(var(--width) * 0.025) 0
      hsl(210 10% 10% / 0.65) inset;
  border-radius: 100vh;
}
input:not(:checked) + button {
  background-color: #3c86bc;
}
input:checked + button {
  background-color: #1e223c;
}
button .🌞 {
  position: absolute;
  top: 13px;
  width: 120px;
  height: 120px;
  overflow: hidden;
  border-radius: 50%;
  z-index: 10;
  background-color: rgb(243, 203, 43);
  box-shadow: 0 5px 5px 0 rgb(215, 222, 234) inset, 0 -5px 5px 0 rgb(175, 156, 80) inset,
    0 5px 6px 0 rgb(109, 106, 98);
}
button .🌞 > .🌛 {
  position: absolute;
  left: 120px;
  width: 120px;
  height: 120px;
  border-radius: 50%;
  z-index: 10;
  background-color: rgb(202, 207, 214);
  box-shadow: 0 5px 5px 0 rgb(215, 222, 234) inset, 0 -5px 5px 0 rgb(131, 131, 131) inset,
    0 5px 6px 0 rgb(109, 106, 98);
}
button .🌞 > .🌛 > .🌕 {
  position: absolute;
  top: 50px;
  left: 20px;
  width: 40px;
  height: 40px;
  overflow: hidden;
  border-radius: 50%;
  background-color: rgb(151, 157, 176);
  box-shadow: 0 3px 1px 0 #7d7d7d inset, 0 -3px 1px 0 #ababab inset;
}
button .🌞 > .🌛 > .🌕:nth-child(2) {
  top: 20px;
  left: 50px;
  width: 20px;
  height: 20px;
}
button .🌞 > .🌛 > .🌕:nth-child(3) {
  top: 72px;
  left: 80px;
  width: 22px;
  height: 22px;
}
input:not(:checked) + button > .🌞 {
  transform: translate(0, 0);
  transition: all 400ms linear;
}
input:not(:checked):hover + button > .🌞 {
  transform: translateX(20px);
}
input:checked + button > .🌞 {
  transform: translate(246px, 0);
  transition: all 400ms linear;
}
input:not(:checked) + button > .🌞 .🌛 {
  transform: translate(0, 0);
  transition: all 400ms linear;
}
input:checked:hover + button > .🌞 {
  transform: translateX(226px);
}
input:checked + button > .🌞 .🌛 {
  left: 0;
  transform: translate(0, 0);
  transition: all 400ms linear;
}
button > .⭕️ {
  position: absolute;
  height: 220px;
  width: 220px;
  top: -40px;
  left: -40px;
  border-radius: 50%;
  transition: all 400ms linear;
}
button > .⭕️::before {
  position: absolute;
  content: "";
  top: -50px;
  left: -50px;
  height: 320px;
  width: 320px;
  border-radius: 50%;
  transition: all 400ms linear;
}
button > .⭕️::after {
  position: absolute;
  content: "";
  top: -100px;
  left: -100px;
  height: 420px;
  width: 420px;
  border-radius: 50%;
  transition: all 400ms linear;
}
input:not(:checked) + button > .⭕️ {
  background: #ffffff0d;
}
input:not(:checked) + button > .⭕️::before {
  background: #ffffff0d;
}
input:not(:checked) + button > .⭕️::after {
  background: #ffffff0d;
}
input:checked + button > .⭕️ {
  left: 210px;
  background-color: rgba(120, 120, 120, 0.2);
}
input:checked + button > .⭕️::before {
  background-color: rgba(120, 120, 120, 0.2);
}
input:checked + button > .⭕️::after {
  background-color: rgba(120, 120, 120, 0.2);
}
input:not(:checked):hover + button > .⭕️ {
  transform: translateX(5%);
}
input:not(:checked):hover + button > .⭕️::before {
  transform: translateX(5%);
}
input:not(:checked):hover + button > .⭕️::after {
  transform: translateX(10%);
}
input:checked:hover + button > .⭕️ {
  transform: translateX(-5%);
}
input:checked:hover + button > .⭕️::before {
  transform: translateX(-5%);
}
input:checked:hover + button > .⭕️::after {
  transform: translateX(-10%);
}
.💭 {
  filter: drop-shadow(-20px -20px 0 rgba(255, 255, 255, 0.4));
  position: absolute;
  top: 0;
  right: 0;
}
.💭 .☁️ {
  position: absolute;
  width: 80px;
  height: 80px;
  background-color: #fff;
  border-radius: 50%;
  box-shadow: 0 0 6px 0 #fff;
}
.💭 .☁️:nth-child(1) {
  top: 20px;
  right: -20px;
}
.💭 .☁️:nth-child(2) {
  top: 60px;
  right: 0;
}
.💭 .☁️:nth-child(3) {
  top: 100px;
  right: 40px;
}
.💭 .☁️:nth-child(4) {
  top: 120px;
  right: 80px;
}
.💭 .☁️:nth-child(5) {
  top: 110px;
  right: 130px;
}
.💭 .☁️:nth-child(6) {
  top: 100px;
  right: 190px;
}
.💭 .☁️:nth-child(7) {
  top: 110px;
  right: 230px;
}
.💭 .☁️:nth-child(8) {
  top: 120px;
  right: 300px;
}
.✈️ {
  position: absolute;
  bottom: -50px;
  right: -100px;
  width: 20%;
  transform: rotate(30deg);
}
input:not(:checked) + button > .✈️ {
  opacity: 1;
  transform: translate(-500px, -200px) rotate(0);
  transition: all 10s linear;
}
.🐻 {
  position: absolute;
  bottom: -80px;
  left: -50px;
  width: 10%;
  rotate: calc(360deg);
}
input:checked + button > .🐻 {
  transform: translate(300px, -250px) rotate(-720deg);
  transition: all 10s linear;
}
input:checked + button > .💭 {
  animation: fadeOut 600ms linear 0s 1 forwards reverse;
}
input:checked + button > .💭 {
  animation: fadeOut 600ms linear 0s 1 forwards reverse;
}
input:not(:checked) + button > .💭 {
  animation: fadeIn 600ms linear 0s 1 forwards reverse;
}
@keyframes fadeIn {
  0% {
    transform: translateY(0);
  }
  20% {
    transform: translateY(-5px);
  }
  100% {
    transform: translateY(240px);
    filter: drop-shadow(-20px 120px 0 rgba(255, 255, 255, 0.4));
  }
}
@keyframes fadeOut {
  0% {
    transform: translateY(140px);
  }
  40% {
    transform: translateY(120px);
  }
  100% {
    transform: translateY(0);
    filter: drop-shadow(-20px 120px 0 rgba(255, 255, 255, 0.4));
  }
}
.🌟 {
  position: absolute;
  top: 0;
  transition: all 300ms linear;
}
@keyframes twinkle {
  0%,
  40%,
  60%,
  100% {
    transform: scale(0.8);
  }
  50% {
    transform: scale(0);
  }
}
@keyframes twinkle-4 {
  0%,
  40%,
  60%,
  100% {
    transform: scale(0.4);
  }
  50% {
    transform: scale(0);
  }
}
.🌟 .✨ {
  position: absolute;
  top: 20px;
  left: 60px;
  width: 20px;
  height: 20px;
  background-image: radial-gradient(transparent 70%, #fff 100%);
  background-size: 20px;
  background-repeat: repeat;
  background-position: 10px 10px;
  transition: all 400ms linear;
}
.🌟 .✨:nth-child(2) {
  top: 44px;
  left: 34px;
  transform: scale(0.6);
  animation: twinkle 4s -8s infinite;
}
.🌟 .✨:nth-child(3) {
  top: 56px;
  left: 64px;
  transform: scale(0.6);
}
.🌟 .✨:nth-child(4) {
  top: 96px;
  left: 40px;
  transform: scale(0.3);
}
.🌟 .✨:nth-child(5) {
  top: 86px;
  left: 50px;
  transform: scale(0.3);
}
.🌟 .✨:nth-child(6) {
  top: 106px;
  left: 70px;
  transform: scale(0.4);
  animation: twinkle-4 5s -10s infinite;
}
.🌟 .✨:nth-child(7) {
  top: 36px;
  left: 166px;
  transform: scale(0.45);
}
.🌟 .✨:nth-child(8) {
  top: 56px;
  left: 156px;
  transform: scale(0.3);
}
.🌟 .✨:nth-child(9) {
  top: 40px;
  left: 202px;
  transform: scale(1.2);
}
.🌟 .✨:nth-child(10) {
  top: 70px;
  left: 192px;
  transform: scale(0.4);
  animation: twinkle-4 4s -2s infinite;
}
.🌟 .✨:nth-child(11) {
  top: 90px;
  left: 172px;
  transform: scale(0.8);
  animation: twinkle 3s -3s infinite;
}
input:not(:checked) + button > .🌟 {
  animation: starFadeOut 600ms linear 0s 1 forwards;
}
input:checked + button > .🌟 {
  animation: starFadeIn 600ms linear 0s 1 forwards;
}
@keyframes starFadeIn {
  0% {
    transform: translateY(-30px);
  }
  20% {
    transform: translateY(10px);
    opacity: 0.6;
  }
  80% {
    opacity: 0.8;
  }
  100% {
    transform: translateY(0);
    opacity: 1;
  }
}
@keyframes starFadeOut {
  0% {
    transform: translateY(0);
    opacity: 1;
  }
  20% {
    transform: translateY(-5px);
  }
  100% {
    transform: translateY(-240px);
    opacity: 0;
  }
}
</style>
