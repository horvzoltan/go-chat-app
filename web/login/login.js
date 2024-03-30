document.addEventListener("DOMContentLoaded", function () {
    const loginForm = document.getElementById("loginForm");

    loginForm.addEventListener("submit", function (event) {
        event.preventDefault();

        const formData = new FormData(loginForm);
        const data = {
            username: formData.get("username"),
            password: formData.get("password"),
        };

        fetch("/api/login", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
            },
            body: JSON.stringify(data),
        })
            .then(response => {
                if (!response.ok) {
                    throw new Error('Login failed');
                }
                return response.text();
            })
            .then(result => {
                alert("Login successful!");
            })
            .catch(error => {
                console.error("Error:", error);
                alert("Login failed. Please check your username and password.");
            });
    });
});

