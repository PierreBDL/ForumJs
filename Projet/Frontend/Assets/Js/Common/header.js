const bodyDom = document.body;
const headerDom = document.createElement("header");

const header = (pageName) => {

    headerDom.innerHTML = "";
    
    // Logo
    let logoLink = document.createElement("a");
    logoLink.href = "/home";
    let logo = document.createElement("img");
    logo.src = "/Assets/Images/logo.png";
    logo.alt = "Logo";
    logo.id = "logo";
    logoLink.appendChild(logo);
    headerDom.appendChild(logoLink);

    // Nav
    let nav = document.createElement("nav");

    // Liens
    let link = document.createElement("a");
    link.href = "/home";
    link.textContent = "Accueil";
    if (link.textContent.toLocaleUpperCase() === pageName.toLocaleUpperCase()) {
        link.classList.add("active");
    }
    nav.appendChild(link);

    let link2 = document.createElement("a");
    link2.href = "/discussions";
    link2.textContent = "Discussions";
    if (link2.textContent.toLocaleUpperCase() === pageName.toLocaleUpperCase()) {
        link2.classList.add("active");
    }
    nav.appendChild(link2);

    // Nav dans header
    headerDom.appendChild(nav);

    // Page profil
    let profilLink = document.createElement("a");
    profilLink.href = "/profil";
    let profil = document.createElement("img");
    profil.src = localStorage.getItem("profilePhoto") || "/Assets/Images/profil.jpg";
    profil.alt = "Profil";
    profil.id = "profil";
    profilLink.appendChild(profil);
    headerDom.appendChild(profilLink);

    // Header dans body
    bodyDom.appendChild(headerDom);
}

const dataPage = document.body.getAttribute("data-page");
header(dataPage);