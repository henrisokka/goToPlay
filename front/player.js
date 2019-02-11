var Tone = require('tone');

console.log("player playing playable plays");

function testMusic(argument, socket) {
    console.log("testing music with argument", argument);
    const synth = new Tone.MonoSynth();
}

module.exports = {
    testMusic: testMusic,
}