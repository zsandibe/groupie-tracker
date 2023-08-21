// Function to update value display element for a slider
function updateSliderValueDisplay(sliderId, valueDisplayId) {
    const slider = document.getElementById(sliderId);
    const valueDisplay = document.getElementById(valueDisplayId);
    valueDisplay.textContent = slider.value;
}

// Add event listeners for slider controls
document.addEventListener("DOMContentLoaded", () => {
    updateSliderValueDisplay("fromSlider1", "fromSlider1Value");
    updateSliderValueDisplay("toSlider1", "toSlider1Value");
    updateSliderValueDisplay("fromSlider2", "fromSlider2Value");
    updateSliderValueDisplay("toSlider2", "toSlider2Value");

    document.getElementById("fromSlider1").addEventListener("input", () => {
        updateSliderValueDisplay("fromSlider1", "fromSlider1Value");
    });

    document.getElementById("toSlider1").addEventListener("input", () => {
        updateSliderValueDisplay("toSlider1", "toSlider1Value");
    });

    document.getElementById("fromSlider2").addEventListener("input", () => {
        updateSliderValueDisplay("fromSlider2", "fromSlider2Value");
    });

    document.getElementById("toSlider2").addEventListener("input", () => {
        updateSliderValueDisplay("toSlider2", "toSlider2Value");
    });

    // Add similar listeners for other sliders if needed
});
