 const menuButton = document.getElementById("menu-button");
const sideMenu = document.getElementById("side-menu");
const closeButton = document.getElementById("close-button");

menuButton.addEventListener("click", () => {
    sideMenu.classList.add("active");
    menuButton.style.visibility = "hidden"; // Скрываем кнопку при открытии меню
});

closeButton.addEventListener("click", () => {
    sideMenu.classList.remove("active");
    menuButton.style.visibility = "visible"; // Показываем кнопку при закрытии меню
});
const deleteFiltersButton = document.getElementById("delete-filters");     
  deleteFiltersButton.addEventListener("click", () => {
    // Сбрасываем значения всех полей фильтрации
    document.getElementById("fromSlider1").value = 1900;
    document.getElementById("toSlider1").value = 2022;
    document.getElementById("fromInput1").value = 1900;
    document.getElementById("toInput1").value = 2022;

    document.getElementById("fromSlider2").value = 1900;
    document.getElementById("toSlider2").value = 2022;
    document.getElementById("fromInput2").value = 1900;
    document.getElementById("toInput2").value = 2022;

    const membersCheckboxes = document.querySelectorAll(".form-check-input[name='memberscount']");
    membersCheckboxes.forEach(checkbox => {
        checkbox.checked = false;
    });

    // Очищаем поле для мест
    document.querySelector("input[name='location']").value = "";

    // Если нужно, добавьте другие поля фильтрации для сброса

    // Закрываем боковое меню
    sideMenu.classList.remove("active");
    menuButton.style.visibility = "visible";
  });