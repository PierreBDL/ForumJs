document.getElementById("submitBtn")?.addEventListener("click", (e) => {
    const username = document.getElementById("username").value;
    const email = document.getElementById("email").value;
    const password = document.getElementById("password").value;
    const messageDiv = document.getElementById("message");
    const forbidenCaraters = /[`!@#$%^&*()_+\-=\[\]{};':"\\|,.<>\/?~]/;



    if (username === "" || email === "" || password === "") {
        message.style.display = "flex";
        messageDiv.textContent = "Veuillez remplir tout les champs";
        messageDiv.style.color = "red";
        return;
    } else {
        if (password.length < 8) {
            message.style.display = "flex";
            messageDiv.textContent = "Le mot de passe doit contenir au moins 8 caractères";
            messageDiv.style.color = "red";
            return;
        } else if (username.length > 20) {
            message.style.display = "flex";
            messageDiv.textContent = "Le nom d'utilisateur ne doit pas dépasser 10 caractères";
            messageDiv.style.color = "red";
            return;
        } else if (forbidenCaraters.test(username)) {
            message.style.display = "flex";
            messageDiv.textContent = "Le nom d'utilisateur ne doit pas contenir de caractères spéciaux";
            messageDiv.style.color = "red";
            return;
        }

        register();
    }
})

async function register() {
    const username = document.getElementById("username").value;
    const email = document.getElementById("email").value;
    const password = document.getElementById("password").value;
    const messageDiv = document.getElementById("message");

    try {
        let result = await fetch("/api/register", {
            method: "POST",
            headers: {
                "Content-Type": "application/json"
            },
            body: JSON.stringify({
                username: username,
                email: email,
                password: password,
                avatar: "/Assets/Images/profil.jpg",
                createAt: new Date().toISOString(),
                lastConnexion: new Date().toISOString(),
            })
        });

        let data = await result.json();

        if (result.status === 200) {
            message.style.display = "flex";
            messageDiv.textContent = "Inscription réussie";
            messageDiv.style.backgroundColor = "rgb(122, 216, 122)";
            messageDiv.style.color = "green";

            // Enregistrer dans le local storage
            localStorage.setItem("username", username);
            localStorage.setItem("email", email);
            localStorage.setItem("profilePhoto", "/Assets/Images/profil.jpg");
            location.reload();
            return;
        } else if (data.message != "") {
            message.style.display = "flex";
            messageDiv.textContent = data.message;
            messageDiv.style.color = "red";
        }
    } catch (error) {
        console.log(error);
    }
}