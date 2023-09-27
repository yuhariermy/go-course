var texts = ['2 Jam Sampai', 'Order Online', 'Pengiriman 24 Jam', 'Order 2 Menit', 'Bisa Custom', 'Dapat Cashback', 'Kirim Bunga Papan', 'Kirim Bunga Buket'];


// Function to generate and display random text
function generateRandomText() {
    var randomIndex = Math.floor(Math.random() * texts.length);
    var randomText = texts[randomIndex];
    var randomTextElement = document.getElementById("randomText");

    randomTextElement.style.opacity = 0;
    randomTextElement.style.transform = "translateY(100%)";

    requestAnimationFrame(function () {
        randomTextElement.textContent = randomText;
        randomTextElement.style.opacity = 1;
        randomTextElement.style.transform = "translateY(0)";
    });
}
// Change the text every 2 seconds
setInterval(generateRandomText, 4000);
// setInterval(generateRandomText, getAnimationTime());

// Initial text generation
generateRandomText();