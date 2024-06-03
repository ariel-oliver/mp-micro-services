
async function handleLogin(event) {
    event.preventDefault(); 

    const username = document.getElementById('username').value;
    const password = document.getElementById('password').value;

    try {
        const response = await fetch('http://localhost:8080/api/auth/login', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({ username, password }),
        });

        if (!response.ok) {
            throw new Error('Login failed');
        }

        const data = await response.json();
        console.log('Login successful:', data);
        
    } catch (error) {
        console.error('Error:', error);
        document.getElementById('login-error').style.display = 'block';
    }
}


document.getElementById('login-form').addEventListener('submit', handleLogin);


async function handleRegister(event) {
    event.preventDefault(); 

    const username = document.getElementById('username').value;
    const password = document.getElementById('password').value;

    try {
        const response = await fetch('http://localhost:8080/api/auth/register', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({ username, password }),
        });

        if (!response.ok) {
            throw new Error('Registration failed');
        }

        const data = await response.json();
        console.log('Registration successful:', data);
        
    } catch (error) {
        console.error('Error:', error);
        document.getElementById('register-error').style.display = 'block';
    }
}


document.getElementById('register-form').addEventListener('submit', handleRegister);


async function fetchProducts() {
    try {
        const response = await fetch('http://localhost:8080/api/products');
        if (!response.ok) {
            throw new Error('Failed to fetch products');
        }
        const products = await response.json();
        const productsList = document.getElementById('products-list');
        products.forEach(product => {
            const productCard = document.createElement('div');
            productCard.className = 'column is-one-quarter';
            productCard.innerHTML = `
                <div class="card">
                    <div class="card-image">
                        <figure class="image is-4by3">
                            <img src="${product.image}" alt="${product.name}">
                        </figure>
                    </div>
                    <div class="card-content">
                        <p class="title">${product.name}</p>
                        <p class="subtitle">$${product.price}</p>
                    </div>
                    <div class="card-footer">
                        <a href="#" class="card-footer-item">Add to cart</a>
                    </div>
                </div>
            `;
            productsList.appendChild(productCard);
        });
    } catch (error) {
        console.error('Error:', error);
    }
}


document.addEventListener('DOMContentLoaded', fetchProducts)