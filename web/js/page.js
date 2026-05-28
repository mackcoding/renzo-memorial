document.addEventListener("DOMContentLoaded", function () {
  const modal = document.getElementById("photoModal");
  const modalImg = document.getElementById("modalImage");
  const modalCaption = document.getElementById("modalCaption");
  const closeBtn = document.querySelector(".photo-modal-close");

  document.addEventListener("click", function (e) {
    if (e.target.matches(".card img, .memorial-card img")) {
      const img = e.target;
      const card = img.closest(".card, .memorial-card");

      if (!card) return;

      modal.classList.add("show");
      modalImg.src = img.src;
      modalImg.alt = img.alt;

      const captionEl = card.querySelector(".card-text, .card-body p");
      modalCaption.textContent = captionEl ? captionEl.textContent : img.alt;
    }
  });

  closeBtn.addEventListener("click", function () {
    modal.classList.remove("show");
  });

  modal.addEventListener("click", function (e) {
    if (e.target === modal) {
      modal.classList.remove("show");
    }
  });

  document.addEventListener("keydown", function (e) {
    if (e.key === "Escape" && modal.classList.contains("show")) {
      modal.classList.remove("show");
    }
  });
});