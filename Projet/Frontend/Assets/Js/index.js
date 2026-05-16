async function getPosts() {
    try {
        let result = await fetch("/api/getPosts", {
            method: "GET",
            headers: {
                "Content-Type": "application/json"
            }
        })

        if (!result.ok) {
            return null;
        }

        let data = await result.json();
        return data;

    } catch (error) {
        console.log(error);
    }
}

function displayPosts(posts) {
    const container = document.getElementById("feedContainer");
    container.innerHTML = "";

    if (!posts || posts.length === 0) {
        container.textContent = "Aucun post.";
        return;
    }

    posts.forEach(post => {
        const authorName = post.User ? post.User.Username : "Anonyme";

        container.innerHTML += `
            <div class="card">
                <div class="headerPost">
                    <span class="postAuthor">@<a href="/profilVisit?username=` + authorName + `">` + authorName + `</a><span>
                    <h3 class="postTitle">` + post.Title + `</h3>
                </div>
                <div class="contentPost">
                    <p class="postContent">` + post.Content + `</p>
                </div>
                <div class="footerPost">
                    <span class="postDate">` + new Date(post.CreatedAt).toLocaleDateString() + `</span>
                    <button class="view-btn" data-id="` + post.ID + `">Voir le post</button>
                </div>
            </div>
            `;
    });
}

document.addEventListener("DOMContentLoaded", () => {
    getPosts().then(posts => {
        displayPosts(posts);
    });
});