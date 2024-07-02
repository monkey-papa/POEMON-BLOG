<template>
  <div class="letter-wrap">
    <div id="jsi-cherry-container" style="z-index: -99">
      <canvas width="646" height="754" position="absolute" top="0"></canvas>
    </div>

    <section class="container" id="contact">
      <form class="flip">
        <div class="front">
          <h2>To&nbsp;&nbsp;Ming</h2>
          <a id="flip2back" href="#">Flip</a>
        </div>
        <div id="content" class="back">
          <div id="letter">
            <div class="container">
              <div class="flip">
                <div class="front"></div>
                <div class="back">
                  <p class="letter">
                    Ming<br /><br />&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;夜にはいつも寒いよね、でも、手を繋いでいると、暖かくなるよ！<br />&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;どんなに寒い夜も、君と二人でいれば、ちっとも寒くない！<br /><br />
                  </p>
                  <p style="float: right; display: block; width: 80px">Hao</p>
                  <span class="typed-cursor" style="animation-iteration-count: infinite">|</span>
                  <a id="close" href="#">Close</a>
                </div>
              </div>
            </div>
          </div>
          <div id="top">
            <a id="flip" href="#contact">Flip</a>
          </div>
          <div id="lid" class="container">
            <div class="flip">
              <div class="front">
                <a id="open" href="#content"><span style="background-position: 0px -150px"></span></a>
              </div>
              <div class="back"></div>
            </div>
          </div>
        </div>
      </form>
    </section>

    <div id="footer">
      <audio id="music2" preload="auto" loop="" src=""></audio>
      <a id="music_btn2" href="javascript:playPause();" style="opacity: 1"></a>
    </div>
  </div>
</template>

<script>
export default {
  data() {
    return {}
  },
  created() {},
  methods: {}
}
</script>

<style scoped>
.letter-wrap {
  font-family: Satisfy, serif;
}

#jsi-cherry-container {
  position: absolute;
  top: 0;
  width: 100%;
  height: 100%;
  margin: 0;
  padding: 0;
  background-color: transparent;
  overflow: hidden;
  opacity: 0.618;
  z-index: 99;
}

#footer {
  /*width: 100%; */
  position: fixed;
  /*bottom: 100px;*/
  top: 10px;
  right: 10px;
}

#music_btn2 {
  position: relative;
  margin: 0 auto;
  display: block;
  width: 32px;
  height: 32px;
  background: url(https://haiyong.site/demo/eluvletter/images/play.png) no-repeat;
  opacity: 0.2;
}

body {
  opacity: 0;
}

span {
  display: block;
  width: 150px;
  height: 150px;
  background: url(https://haiyong.site/demo/eluvletter/images/stitch.png) no-repeat;
  background-position: 0 0;
}

.typed-cursor {
  opacity: 0;
  -webkit-animation: blink 0.7s infinite;
  -moz-animation: blink 0.7s infinite;
  animation: blink 0.7s infinite;
}

@keyframes blink {
  0% {
    opacity: 0;
  }
  50% {
    opacity: 0;
  }
  100% {
    opacity: 0;
  }
}

@-webkit-keyframes blink {
  0% {
    opacity: 0;
  }
  50% {
    opacity: 0;
  }
  100% {
    opacity: 0;
  }
}

@-moz-keyframes blink {
  0% {
    opacity: 0;
  }
  50% {
    opacity: 0;
  }
  100% {
    opacity: 0;
  }
}

/* reset */
html,
body,
div,
span,
applet,
object,
iframe,
h1,
h2,
h3,
h4,
h5,
h6,
p,
blockquote,
pre,
a,
abbr,
acronym,
address,
big,
cite,
code,
del,
dfn,
em,
img,
ins,
kbd,
q,
s,
samp,
small,
strike,
strong,
sub,
sup,
tt,
var,
b,
u,
i,
center,
dl,
dt,
dd,
ol,
ul,
li,
fieldset,
form,
label,
legend,
table,
caption,
tbody,
tfoot,
thead,
tr,
th,
td,
article,
aside,
canvas,
details,
embed,
figure,
figcaption,
footer,
header,
hgroup,
menu,
nav,
output,
ruby,
section,
summary,
time,
mark,
audio,
video {
  margin: 0;
  padding: 0;
  border: 0;
  font: inherit;
  font-size: 100%;
  vertical-align: baseline;
}

html {
  line-height: 1;
  content-zooming: none;
}

article,
aside,
details,
figcaption,
figure,
footer,
header,
hgroup,
menu,
nav,
section,
summary {
  display: block;
}

/* body */

body {
  font-family: Satisfy, serif;
  color: #837362;
  background: #fbf8f7;
}

h2 {
  font-size: 1.5em;
  line-height: 1.6;
  float: left;
  font-family: Satisfy, serif;
}

/* links */

a {
  color: #9b4849;
  text-decoration: none;
  display: block;
}

a:hover {
  color: #783839;
}

#open {
  width: 150px;
  text-align: center;
  margin-top: 50px;
  margin-left: 200px;
}

#flip {
  margin: 160px 20px 0 0;
  float: right;
}

#flip2back {
  margin-top: 150px;
  float: right;
}

#close {
  position: absolute;
  bottom: 10px;
  right: 10px;
}

/*****************************/

section.container {
  position: relative;
  width: 600px;
  height: 300px;
  margin: 280px auto 0;
  perspective: 1000px;
  -webkit-perspective: 1000px;
  -moz-perspective: 1000px;
}

section.container > form.flip {
  width: 100%;
  height: 100%;
  position: absolute;
  -webkit-transform: rotateY(180deg);
  -moz-transform: rotateY(180deg);
  -webkit-transform-style: preserve-3d;
  -moz-transform-style: preserve-3d;
  -webkit-transition: -webkit-transform 0.7s 0s;
  -moz-transition: -moz-transform 0.7s 0s;
  -webkit-transform-origin: center center;
  -moz-transform-origin: center center;
}

section.container > form.flip > .front,
section.container > form.flip > .back {
  position: absolute;
  width: 100%;
  height: 100%;
  backface-visibility: hidden;
  -webkit-backface-visibility: hidden;
  -moz-backface-visibility: hidden;
}

section.container > form.flip > .back {
  -webkit-transform: rotateY(180deg);
  -moz-transform: rotateY(180deg);
  transform: rotateY(180deg);
  background: #837362;
  box-shadow: inset 0 10px 30px -10px rgba(0, 0, 0, 0.3), inset 0 1px 0 rgba(0, 0, 0, 0.2), inset 0 2px 0 rgba(255, 255, 255, 0.2);
}

section.container:target > form.flip {
  -webkit-transform: rotateY(0deg);
  -moz-transform: rotateY(0deg);
  transform: rotateY(0deg);
}

.no-csstransforms3d section.container:target > form.flip > .back {
  display: none;
}

.container#lid {
  position: relative;
  width: 100%;
  height: 50%;
  -webkit-perspective: 800px;
  -moz-perspective: 800px;
  perspective: 800px;
  z-index: 2;
  transition: z-index 0s 1s;
  -moz-transition: z-index 0s 1s;
  -webkit-transition: z-index 0s 1s;
}

.no-csstransforms3d .container#lid {
  -webkit-transition: none;
  -moz-transition: none;
  transition: none;
}

#content:target .container#lid {
  z-index: -1;
  transition: z-index 0s 0.5s;
  -moz-transition: z-index 0s 0.5s;
  -webkit-transition: z-index 0s 0.5s;
}

.no-csstransforms3d #content:target .container#lid {
  -webkit-transition: none;
  -moz-transition: none;
  transition: none;
}

.no-csstransforms3d .container#lid .back {
  display: none;
}

.no-csstransforms3d #content:target .container#lid .back {
  display: block;
  top: -150px;
}

#lid .flip {
  width: 100%;
  height: 100%;
}

#lid .flip {
  position: absolute;
  -webkit-transform-style: preserve-3d;
  -moz-transform-style: preserve-3d;
  transform-style: preserve-3d;

  -webkit-transition: -webkit-transform 0.5s 1s;
  -moz-transition: -moz-transform 0.5s 1s;
  transition: -moz-transform 0.5s 1s;

  -webkit-transform-origin: left top;
  -moz-transform-origin: left top;
  transform-origin: left top;
}

#lid .flip > .front,
#lid .flip > .back {
  position: absolute;
  width: 100%;
  height: 100%;
  -webkit-backface-visibility: hidden;
  -moz-backface-visibility: hidden;
  backface-visibility: hidden;
}

#lid .flip > .back {
  -webkit-transform: rotateX(180deg);
  -moz-transform: rotateX(180deg);
  transform: rotateX(180deg);
}

#content:target #lid .flip {
  -webkit-transform: rotateX(180deg);
  -moz-transform: rotateX(180deg);
  transform: rotateX(180deg);

  -webkit-transition-duration: 0.5s;
  -moz-transition-duration: 0.5s;
  transition-duration: 0.5s;

  -webkit-transition-delay: 0s;
  -moz-transition-delay: 0s;
  transition-delay: 0s;
}

#letter .flip > div,
#lid .flip > div,
section.container > form.flip > div {
  -webkit-box-sizing: border-box;
  -moz-box-sizing: border-box;
  box-sizing: border-box;
}

#lid .flip .front,
section.container > form.flip .front {
  background: beige url('https://haiyong.site/demo/eluvletter/images/bg.png');
  border: 1px #eae1d5 solid;
  box-shadow: inset 0 0 10px 1px rgba(255, 255, 255, 0.6), 0 2px 3px -2px rgba(0, 0, 0, 0.6);
  padding: 20px 30px;
  color: #837362;
  text-shadow: 0 1px 0 #fff, 0 1px 0 #fff;
}

section.container > form.flip > .front {
  font-size: 1.25em;
  background: beige url('https://haiyong.site/demo/eluvletter/images/bg.png');
  background: url('https://haiyong.site/demo/eluvletter/images/bg.png') no-repeat 450px 20px, url('https://haiyong.site/demo/eluvletter/images/bg.png');
  padding: 120px 20px 20px 40%;
  border: 1px #eae1d5 solid;
}

#lid .flip .front {
  font-size: 1.25em;
  border-radius: 0 0 40px 40px;
}

#lid .flip .back {
  background: #837362;
  border-radius: 40px 40px 0 0;
  border: 35px solid rgba(255, 255, 255, 0.1);
  border-bottom: none;
  box-shadow: inset 0 10px 30px 10px rgba(0, 0, 0, 0.1);
}

#top {
  text-align: right;
  background: beige url('https://haiyong.site/demo/eluvletter/images/bg.png');
  background: -moz-linear-gradient(top, rgba(255, 255, 255, 0) 0%, rgba(163, 140, 130, 0.17) 100%), url('https://haiyong.site/demo/eluvletter/images/bg.png');
  background: -webkit-gradient(linear, left top, left bottom, color-stop(0%, rgba(255, 255, 255, 0)), color-stop(100%, rgba(163, 140, 130, 0.17))), url('https://haiyong.site/demo/eluvletter/images/bg.png');
  background: -webkit-linear-gradient(top, rgba(255, 255, 255, 0) 0%, rgba(163, 140, 130, 0.17) 100%), url('https://haiyong.site/demo/eluvletter/images/bg.png');
  background: -o-linear-gradient(top, rgba(255, 255, 255, 0) 0%, rgba(163, 140, 130, 0.17) 100%), url('https://haiyong.site/demo/eluvletter/images/bg.png');
  background: -ms-linear-gradient(top, rgba(255, 255, 255, 0) 0%, rgba(163, 140, 130, 0.17) 100%), url('https://haiyong.site/demo/eluvletter/images/bg.png');
  -webkit-box-sizing: border-box;
  -moz-box-sizing: border-box;
  box-sizing: border-box;
  width: 100%;
  height: 65%;
  position: absolute;
  left: 0;
  bottom: 0;
  z-index: 1;
  border: 1px #eae1d5 solid;
  box-shadow: inset 0 0 10px 1px rgba(255, 255, 255, 0.6), 0 -2px 4px -3px rgba(0, 0, 0, 0.75), 0 2px 3px -2px rgba(0, 0, 0, 0.6);
}

#letter {
  background: #fafafa;
  width: 90%;
  height: 95%;
  position: absolute;
  left: 5%;
  top: 5%;
  z-index: 0;
  -webkit-transition: 0.5s 0.5s;
  -moz-transition: 0.5s 0.5s;
  -o-transition: 0s 0s;
  transition: 0s 0s;
}

#content:target #letter {
  top: -40%;
  -o-transition: 0.5s 0.5s;
  transition: 0.5s 0.5s;
}

#letter .container {
  position: absolute;
  width: 100%;
  /*height: 50%;*/
  -webkit-perspective: 800px;
  -moz-perspective: 800px;
  perspective: 800px;
  top: 0;
  left: 0;
  -webkit-transition: 0s 0s;
  -moz-transition: 0s 0s;
  transition: 0s 0s;
  z-index: 1;
}

#letter .flip {
  width: 100%;
  height: 100%;
}

#letter .flip {
  position: absolute;
  -webkit-transform-style: preserve-3d;
  -moz-transform-style: preserve-3d;
  transform-style: preserve-3d;

  -webkit-transition: -webkit-transform 0.5s 0s;
  -moz-transition: -moz-transform 0.5s 0s;
  transition: -moz-transform 0.5s 0s;

  -webkit-transform-origin: left top;
  -moz-transform-origin: left top;
  transform-origin: left top;
}

#letter .flip > .front,
#letter .flip > .back {
  position: absolute;
  width: 100%;
  height: 100%;
  -webkit-backface-visibility: hidden;
  -moz-backface-visibility: hidden;
  backface-visibility: hidden;
}

#letter .flip > .back {
  -webkit-transform: rotateX(180deg);
  -moz-transform: rotateX(180deg);
  transform: rotateX(180deg);
}

#content:target #letter .flip {
  -webkit-transform: rotateX(180deg);
  -moz-transform: rotateX(180deg);
  transform: rotateX(180deg);

  -webkit-transition-duration: 0.5s;
  -moz-transition-duration: 0.5s;
  transition-duration: 0.5s;

  -webkit-transition-delay: 1s;
  -moz-transition-delay: 1s;
  transition-delay: 1s;
}

#content:target #letter .container {
  z-index: -1;
  -webkit-transition: z-index 0s 1.5s;
  -moz-transition: z-index 0s 1.5s;
  transition: z-index 0s 1.5s;
}

#letter .flip .front {
  background: #fafafa;
  box-shadow: 0 2px 3px 0 rgba(0, 0, 0, 0.2);
}

.no-csstransforms3d #content:target #letter .flip .front {
  display: none;
}

#letter .flip .back {
  background: #fafafa;
  background: -webkit-linear-gradient(top, #fafafa 90%, #f7f7f7 100%);
  background: -moz-linear-gradient(top, #fafafa 80%, #f7f7f7 100%);
  padding: 40px 40px 0;
  font-size: 18px;
  line-height: 1.4;
  border-bottom: 1px solid var(--maxWhiteMask);
  box-shadow: 0 -5px 8px -4px rgba(0, 0, 0, 0.15);
}

.no-csstransforms3d #content:target #letter .flip .back {
  top: -140px;
}

input[type='text'] {
  background: transparent;
  border: none;
  border-bottom: 1px dotted #837362;
  box-shadow: 0 1px 0 rgba(255, 255, 255, 0.4);
  width: 50%;
  -webkit-appearance: none;
}

input[type='text']:focus {
  outline: none;
  border-bottom: 1px solid #837362;
  -webkit-appearance: none;
}

textarea {
  margin: 10px auto 0;
  -webkit-box-sizing: border-box;
  -moz-box-sizing: border-box;
  box-sizing: border-box;
  width: 90%;
  height: 160px;
  padding: 10px;
  display: block;
  background: transparent;
  background: rgba(255, 255, 255, 0.07);
  border: 1px dashed #c8bfb6;
}

textarea:focus {
  outline: none;
  background-color: rgba(255, 255, 255, 0.12);
  border-color: rgba(255, 255, 255, 0.46);
  box-shadow: inset 0 0 2px rgba(0, 0, 0, 0.1);
}

input[type='submit'] {
  display: block;
  margin: 10px auto;
  -webkit-appearance: none;
  font-size: 0.9em;
  background: #9b4849;
  padding: 5px 20px;
  border-radius: 3px;
  border: 1px solid #9b4849;
  box-shadow: inset 0 1px 0 var(--maxWhiteMask);
  color: #fff;
  -webkit-transition: background 0.5s;
}

input[type='submit']:hover {
  background: #783839;
}

/*-------ie8-----------*/
.container.target > form.flip > .back,
#content.target #letter .flip .front {
  display: none;
}

#content.target .container#lid {
  z-index: -1;
}

#content.target .container#lid .back {
  display: block;
  top: -150px;
}

#content.target #letter {
  top: -40%;
}

#content.target #letter .flip .back {
  top: -140px;
}
</style>
