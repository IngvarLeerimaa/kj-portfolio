function login() {
    // Retrieve the token from localStorage
    let token = localStorage.getItem('token');

    // Check if token is not present or empty
    if (token == undefined || token == null || token == "") {
        // Get form and input elements
        const login = document.getElementById('form');
        const email = document.getElementById('email');
        const password = document.getElementById('password');

        // Add submit event listener to form
        login.addEventListener('submit', async (e) => {
            // Prevent default form submission behavior
            e.preventDefault();
            
            // Encode email and password to base64
            const credentials = btoa(`${email.value}:${password.value}`);
            
            // Send credentials to server for authentication
            const response = await fetch('https://01.kood.tech/api/auth/signin', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                    // Include authorization header with basic authentication credentials
                    Authorization: `Basic ${credentials}`
                }
            });

            // Parse response body as JSON
            let token = await response.json();

            // Check for errors in the response
            if (token.error) {
                // Display error message
                let err = document.getElementById('err');
                err.innerHTML = token.error;
                err.style.display = 'block';
                // Hide error message after 3 seconds
                setTimeout(() => {
                    err.style.display = 'none';
                }, 3000);
            } else {
                // Save token to localStorage
                localStorage.setItem('token', token);
                // Redirect user to main.html
                location.replace('main.html');
            }
        });
    } else {
        // If token exists, redirect user to main.html
        location.replace('main.html');
    }
}

// Call the login function when the script runs
login();
