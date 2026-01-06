// Mobile Navigation Toggle
const navToggle = document.getElementById("navToggle")
const navMenu = document.getElementById("navMenu")

if (navToggle) {
  navToggle.addEventListener("click", () => {
    navMenu.classList.toggle("active")
  })
}

// Close mobile menu when clicking outside
document.addEventListener("click", (e) => {
  if (navMenu && navMenu.classList.contains("active") && !navMenu.contains(e.target) && !navToggle.contains(e.target)) {
    navMenu.classList.remove("active")
  }
})

// Tab Functionality
function showTab(tabName) {
  const tabs = document.querySelectorAll(".tab-content")
  const buttons = document.querySelectorAll(".tab-btn")

  tabs.forEach((tab) => {
    tab.classList.remove("active")
  })

  buttons.forEach((btn) => {
    btn.classList.remove("active")
  })

  const selectedTab = document.getElementById(tabName)
  if (selectedTab) {
    selectedTab.classList.add("active")
  }

  event.target.classList.add("active")
}

// Copy Code Functionality
function copyCode(button) {
  const codeBlock = button.parentElement
  const code = codeBlock.querySelector("code").textContent

  navigator.clipboard
    .writeText(code)
    .then(() => {
      const originalText = button.textContent
      button.textContent = "Copié !"
      button.style.background = "#10b981"

      setTimeout(() => {
        button.textContent = originalText
        button.style.background = ""
      }, 2000)
    })
    .catch((err) => {
      console.error("[v0] Erreur lors de la copie:", err)
      button.textContent = "Erreur"
      setTimeout(() => {
        button.textContent = "Copier"
      }, 2000)
    })
}

// Smooth Scroll for Anchor Links
document.querySelectorAll('a[href^="#"]').forEach((anchor) => {
  anchor.addEventListener("click", function (e) {
    e.preventDefault()
    const target = document.querySelector(this.getAttribute("href"))
    if (target) {
      target.scrollIntoView({
        behavior: "smooth",
        block: "start",
      })
    }
  })
})

// Add scroll effect to navbar
let lastScroll = 0
const navbar = document.querySelector(".navbar")

window.addEventListener("scroll", () => {
  const currentScroll = window.pageYOffset

  if (currentScroll <= 0) {
    navbar.style.boxShadow = "none"
  } else {
    navbar.style.boxShadow = "0 4px 6px rgba(0, 0, 0, 0.1)"
  }

  lastScroll = currentScroll
})
