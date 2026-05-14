const bodyDomFooter = document.body;
const footerDom = document.createElement("footer");

const footer = () => {

    // Copyright
    let copyright = document.createElement("p");
    copyright.textContent = "Forum JS";
    footerDom.appendChild(copyright);


    // Footer dans body
    bodyDomFooter.appendChild(footerDom);
}

footer();