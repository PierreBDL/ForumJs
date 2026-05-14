let currentUserProfile = null;
let currentEditField = null;

const devMode = true;
const profil = {
    username: localStorage.getItem("username") || "TestUser",
    avatar_url: localStorage.getItem("profilePhoto") || "/Assets/Images/profil.jpg",
    role: "Admin",
    id: "12345",
    created_at: "2026-01-15T10:00:00Z",
    online: true,
    post_count: 5,
    comment_count: 12,
    like_count: 8,
    dislike_count: 2,
    email: "test@gmail.com",
    lastPosts: {
        post1: {
            title: "First post",
            content: "This is the first post content. This is the first post content. This is the first post content. This is the first post content. This is the first post content.",
            date: "2023-07-25"
        },
        post2: {
            title: "Second post",
            content: "This is the second post content.",
            date: "2023-07-26"
        },
        post3: {
            title: "Third post",
            content: "This is the third post content.",
            date: "2023-07-27"
        },
        post4: {
            title: "Fourth post",
            content: "This is the fourth post content.",
            date: "2023-07-28"
        },
    }
};


// Print lastPosts
const postContentDom = document.getElementById("postContent");

function printLastPosts() {
    postContentDom.innerHTML = "";

    for (let [key, value] of Object.entries(profil.lastPosts)) {
        let div = document.createElement("div");
        div.className = "cardPost";
        div.id = key;

        let contentText = value.content;

        // Content max length 100 characters
        if (value.content.length > 100) {
            contentText = value.content.substring(0, 100) + "...";
        }

        div.innerHTML = '<h4> ' + value.title + ' </h4> <p> ' + contentText + ' </p> <span class="postDate"> ' + new Date(value.date).toLocaleDateString() + ' </span> <button class="viewPostButton">See post</button>';
        postContentDom.appendChild(div);
    }
}

printLastPosts();


// Display data
function printInfos(profile) {
    if (!profile) {
        return
    };

    // name
    document.getElementById("pseudoProfil").textContent = profile.username;

    // Update avatar
    document.getElementById("profilPhoto").src = profile.avatar_url;
    localStorage.setItem("profilePhoto", profile.avatar_url);

    if (profile.avatar_url && !profile.avatar_url.startsWith('blob:')) {
        localStorage.setItem("profilePhoto", profile.avatar_url);
    }

    // Online
    const onlineBall = document.getElementById("isOnlineBall");
    const onlineText = document.getElementById("isOnline");
    if (profile.online) {
        onlineBall.classList.add("Online");
        onlineBall.classList.remove("Offline");
        onlineText.textContent = "Online";
        onlineText.style.color = "green";
    } else {
        onlineBall.classList.add("Offline");
        onlineBall.classList.remove("Online");
        onlineText.textContent = "Offline";
        onlineText.style.color = "red";
    }


    // Role
    document.getElementById("userRole").textContent = profile.role || "User";

    // Date
    const creationDateDom = document.getElementById("creationDate");
    let createdDate = new Date(profile.created_at).toLocaleDateString();
    creationDateDom.textContent = createdDate;

    // Connexion Service
    document.getElementById("connexionService").textContent = profil.email;

    // Stats
    document.getElementById("statsPosts").textContent = profile.post_count || 0;
    document.getElementById("statsComments").textContent = profile.comment_count || 0;
    document.getElementById("statsLikesGiven").textContent = profile.like_count || 0;
    document.getElementById("statsDislikesGiver").textContent = profile.dislike_count || 0;
}

// Edit modale
function openEditFieldModal(fieldType) {
    currentEditField = fieldType;

    if (fieldType === 'username') {
        document.getElementById("btnConnexion").innerHTML = "";
        document.getElementById("editFieldInput").style.display = "block";
        document.getElementById("editFieldLabel").textContent = "Username";
        document.getElementById("editFieldInput").value = document.getElementById("pseudoProfil").textContent;
    }

    if (fieldType === 'emailConnect') {
        document.getElementById("btnConnexion").innerHTML = "";
        document.getElementById("editFieldInput").style.display = "block";
        document.getElementById("editFieldLabel").textContent = "Courriel";
        document.getElementById("editFieldInput").value = profil.email;
    }

    document.getElementById("editFieldModal").style.display = "flex";
    document.getElementById("editFieldModalOverlay").style.display = "block";
    document.getElementById("editFieldInput").focus();
}

function closeEditFieldModal() {
  document.getElementById("editFieldModal").style.display = "none";
  document.getElementById("editFieldModalOverlay").style.display = "none";
  currentEditField = null;

  // Remove class
  const modal = document.getElementById("editFieldModal");
  modal.classList.remove("modal-resize");
}

async function saveEditField() {
  if (!currentEditField) {
    return
  };

  let value = document.getElementById("editFieldInput").value.trim();

  if (!value) {
    alert("You didn't enter a value. Please retry !");
    return;
  }

  // Check mail
  if (currentEditField === "username") {
    profil.username = value;
  }

  if (currentEditField === "emailConnect") {
    profil.email = value;
  }

  printInfos(profil);

  closeEditFieldModal();
}

// Update profile image

document.getElementById("editPhotoInput")?.addEventListener('change', async (e) => {
  const file = e.target.files[0];

  if (!file) {
    return
  };

  // Create URL for image
  let urlImage = URL.createObjectURL(file);

  if (devMode) {
    profil.avatar_url = urlImage;
    printInfos(profil);
    return;
  }

  // Update avatar DOM
  document.getElementById('profilPhoto').src = URL.createObjectURL(file);

  // Update header
  header("Profile");
});

// Click on photo to change
document.querySelector('#profilPhoto')?.addEventListener('click', () => {
  document.getElementById("editPhotoInput").click();
});

// Change username

document.getElementById("editUserBtn")?.addEventListener("click", () =>
  openEditFieldModal('username')
);

// Change address service
document.getElementById("editConnexionServiceBtn")?.addEventListener("click", () =>
  openEditFieldModal('emailConnect')
);


// Modal
document.getElementById("cancelEditField")?.addEventListener("click", closeEditFieldModal);
document.getElementById("editFieldModalOverlay")?.addEventListener("click", closeEditFieldModal);
document.getElementById("saveEditField")?.addEventListener("click", (e) => {
  e.preventDefault();
  saveEditField();
});

document.addEventListener("keydown", (e) => {
  if (e.key === "Escape" && document.getElementById("editFieldModal").style.display === "flex") {
    closeEditFieldModal();
  }
});

document.getElementById("editFieldInput").addEventListener("keydown", (e) => {
  if (e.key === "Enter") {
    saveEditField();
  }
});

// Load
function applySavedData() {
  let savedUsername = localStorage.getItem("username");
  let savedProfilePhoto = localStorage.getItem("profilePhoto");

  if (savedUsername) {
    document.getElementById("pseudoProfil").textContent = savedUsername;
  }

  if (savedProfilePhoto) {
    // If it's a temp link, we load default image
    if (savedProfilePhoto.startsWith('blob:')) {
      localStorage.removeItem("profilePhoto");
    } else {
      let profileImg = document.getElementById("profilPhoto");
      if (profileImg) {
        profileImg.src = savedProfilePhoto;
      }
    }
  }
}

applySavedData();
printInfos(profil);