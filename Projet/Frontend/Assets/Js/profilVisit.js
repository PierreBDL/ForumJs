let currentUserProfile = null;
let currentEditField = null;

const profil = {
  username: localStorage.getItem("username") || "TestUser",
  avatar_url: localStorage.getItem("profilePhoto") || "/Assets/Images/profil.jpg",
  role: "Admin",
  created_at: "2026-01-15T10:00:00Z",
  online: true,
  lastConnexion: "Il y a 1 heure",
  post_count: 5,
  comment_count: 12,
  like_count: 8,
  dislike_count: 2,
  email: "Privé",
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

// Get informations from server with username
async function getProfil() {
  const urlParams = new URLSearchParams(window.location.search);
  const username = urlParams.get("username");

  try {
    let result = await fetch("/api/profil?username=" + username, {
      method: "GET",
      headers: {
        "Content-Type": "application/json"
      }
    });

    let data = await result.json();

    if (result.status === 200) {
      profil.username = data.username;
      profil.avatar_url = data.avatarLink;
      profil.role = data.role;
      profil.created_at = data.createAt;
      profil.online = data.isOnline;
      profil.post_count = data.post_count;
      profil.comment_count = data.comment_count;
      profil.like_count = data.like_count;
      profil.dislike_count = data.dislike_count;
      profil.lastPosts = data.lastPosts;

      // Si en ligne -> dernière connexion = maintenant
      if (profil.online === 1 || profil.online === true) {
        profil.lastConnexion = "Actuellement en ligne"
      } else {
        profil.lastConnexion = data.lastConnexion;
      }

      printInfos(profil);
      printLastPosts();
    } else if (data.message != "") {
      console.log(data.message);
    }
  } catch (error) {
    console.log(error);
  }
}


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
  if (profile.online === 1 || profile.online === true) {
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

  // Last connexion
  document.getElementById("lastVisit").textContent = profile.lastConnexion;

  // Role
  document.getElementById("userRole").textContent = profile.role || "User";

  // Date
  const creationDateDom = document.getElementById("creationDate");
  const createdDate = new Date(profile.created_at).toLocaleDateString();
  creationDateDom.textContent = createdDate;

  // Connexion Service
  document.getElementById("connexionService").textContent = profil.email;

  // Stats
  document.getElementById("statsPosts").textContent = profile.post_count || 0;
  document.getElementById("statsComments").textContent = profile.comment_count || 0;
  document.getElementById("statsLikesGiven").textContent = profile.like_count || 0;
  document.getElementById("statsDislikesGiver").textContent = profile.dislike_count || 0;
}

getProfil();