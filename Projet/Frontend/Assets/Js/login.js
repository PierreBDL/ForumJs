document.getElementById("submitBtn")?.addEventListener("click", (e) => {
    const username = document.getElementById("username").value;
    const password = document.getElementById("password").value;
    const messageDiv = document.getElementById("message");
    const forbidenCaraters = /[`!@#$%^&*()_+\-=\[\]{};':"\\|,.<>\/?~]/;



    if (username === "" || password === "") {
        message.style.display = "flex";
        messageDiv.textContent = "Veuillez remplir tout les champs";
        messageDiv.style.color = "red";
        return;
    } else {
        if (password.length < 8) {
            message.style.display = "flex";
            messageDiv.textContent = "Identifiant ou mot de passe incorect !";
            messageDiv.style.color = "red";
            return;
        } else if (username.length > 20) {
            message.style.display = "flex";
            messageDiv.textContent = "Identifiant ou mot de passe incorect !";
            messageDiv.style.color = "red";
            return;
        } else if (forbidenCaraters.test(username)) {
            message.style.display = "flex";
            messageDiv.textContent = "Identifiant ou mot de passe incorect !";
            messageDiv.style.color = "red";
            return;
        }

        login();
    }
})

async function login() {
    const username = document.getElementById("username").value;
    const password = document.getElementById("password").value;
    const messageDiv = document.getElementById("message");

    try {
        let result = await fetch("/api/login", {
            method: "POST",
            headers: {
                "Content-Type": "application/json"
            },
            body: JSON.stringify({
                username: username,
                password: password,
            })
        });

        let data = await result.json();

        if (result.status === 200) {
            message.style.display = "flex";
            messageDiv.textContent = "Connexion réussie";
            messageDiv.style.backgroundColor = "rgb(122, 216, 122)";
            messageDiv.style.color = "green";
            localStorage.setItem("username", username);
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