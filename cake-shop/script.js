// Cake Data with Prices for Different Sizes
const cakes = [
    {
        id: 1,
        name: "Chocolate Fudge Cake",
        description: "Rich chocolate layers with fudge frosting",
        image: "https://images.unsplash.com/photo-1578985545062-69928b1d9587?w=400",
        prices: {
            6: 35.00,
            8: 50.00,
            10: 70.00,
            12: 95.00
        }
    },
    {
        id: 2,
        name: "Vanilla Bean Cake",
        description: "Classic vanilla with buttercream frosting",
        image: "https://images.unsplash.com/photo-1565958011703-44f9829ba187?w=400",
        prices: {
            6: 30.00,
            8: 45.00,
            10: 65.00,
            12: 85.00
        }
    },
    {
        id: 3,
        name: "Red Velvet Cake",
        description: "Southern classic with cream cheese frosting",
        image: "https://images.unsplash.com/photo-1614707267537-b85aaf00c4b7?w=400",
        prices: {
            6: 40.00,
            8: 55.00,
            10: 75.00,
            12: 100.00
        }
    },
    {
        id: 4,
        name: "Strawberry Shortcake",
        description: "Fresh strawberries with whipped cream",
        image: "https://images.unsplash.com/photo-1565958011703-44f9829ba187?w=400",
        prices: {
            6: 38.00,
            8: 52.00,
            10: 72.00,
            12: 95.00
        }
    },
    {
        id: 5,
        name: "Carrot Cake",
        description: "Spiced carrot cake with cream cheese frosting",
        image: "https://images.unsplash.com/photo-1566940954000-25e29e7bfcae?w=400",
        prices: {
            6: 36.00,
            8: 50.00,
            10: 70.00,
            12: 92.00
        }
    },
    {
        id: 6,
        name: "Lemon Drizzle Cake",
        description: "Zesty lemon with sweet glaze",
        image: "https://images.unsplash.com/photo-1519340333755-56e9c1d04579?w=400",
        prices: {
            6: 32.00,
            8: 47.00,
            10: 67.00,
            12: 88.00
        }
    }
];

// Shopping Cart
let cart = [];

// Initialize the page
document.addEventListener('DOMContentLoaded', () => {
    displayCakes();
    updateCart();
    setMinDate();
});

// Display all cakes
function displayCakes() {
    const container = document.getElementById('cakes-container');
    container.innerHTML = cakes.map(cake => `
        <div class="cake-card" data-id="${cake.id}">
            <img src="${cake.image}" alt="${cake.name}" class="cake-image">
            <div class="cake-info">
                <h3 class="cake-name">${cake.name}</h3>
                <p class="cake-description">${cake.description}</p>
                
                <div class="price-tag">
                    <h4><i class="fas fa-tag"></i> Price by Size:</h4>
                    ${Object.entries(cake.prices).map(([size, price]) => `
                        <div class="size-price">
                            <span class="size-label">${size}" inch</span>
                            <span class="size-cost">$${price.toFixed(2)}</span>
                        </div>
                    `).join('')}
                </div>
                
                <div class="size-selection">
                    <label>Select Size:</label>
                    <div class="size-options">
                        ${Object.keys(cake.prices).map(size => `
                            <div class="size-option">
                                <input type="radio" 
                                       id="cake${cake.id}-size${size}" 
                                       name="cake${cake.id}-size" 
                                       value="${size}"
                                       onchange="updatePrice(${cake.id}, ${size})">
                                <label for="cake${cake.id}-size${size}">${size}"</label>
                            </div>
                        `).join('')}
                    </div>
                </div>
                
                <div class="selected-price" id="price-cake${cake.id}" style="display:none; margin: 1rem 0; padding: 0.5rem; background: #e3f2fd; border-radius: 5px; text-align: center;">
                    <strong>Selected: $<span id="amount-cake${cake.id}">0.00</span></strong>
                </div>
                
                <button class="btn-add-to-cart" onclick="addToCart(${cake.id})">
                    <i class="fas fa-shopping-cart"></i> Add to Cart
                </button>
            </div>
        </div>
    `).join('');
}

// Update displayed price when size is selected
function updatePrice(cakeId, size) {
    const cake = cakes.find(c => c.id === cakeId);
    const price = cake.prices[size];
    
    document.getElementById(`price-cake${cakeId}`).style.display = 'block';
    document.getElementById(`amount-cake${cakeId}`).textContent = price.toFixed(2);
}

// Add item to cart
function addToCart(cakeId) {
    const cake = cakes.find(c => c.id === cakeId);
    const sizeInput = document.querySelector(`input[name="cake${cakeId}-size"]:checked`);
    
    if (!sizeInput) {
        alert('Please select a cake size!');
        return;
    }
    
    const size = sizeInput.value;
    const price = cake.prices[size];
    
    const cartItem = {
        id: Date.now(),
        cakeId: cake.id,
        name: cake.name,
        image: cake.image,
        size: size,
        price: price
    };
    
    cart.push(cartItem);
    updateCart();
    
    // Show success message
    alert(`${cake.name} (${size}") added to cart!`);
    
    // Reset selection
    sizeInput.checked = false;
    document.getElementById(`price-cake${cakeId}`).style.display = 'none';
}

// Update cart display
function updateCart() {
    const cartItems = document.getElementById('cart-items');
    const cartCount = document.getElementById('cart-count');
    
    cartCount.textContent = cart.length;
    
    if (cart.length === 0) {
        cartItems.innerHTML = '<p class="empty-cart">Your cart is empty</p>';
    } else {
        cartItems.innerHTML = cart.map(item => `
            <div class="cart-item">
                <img src="${item.image}" alt="${item.name}" class="cart-item-image">
                <div class="cart-item-details">
                    <h4>${item.name}</h4>
                    <p>Size: ${item.size}" inch</p>
                </div>
                <div class="cart-item-price">
                    <h3>$${item.price.toFixed(2)}</h3>
                    <button class="btn-remove" onclick="removeFromCart(${item.id})">
                        <i class="fas fa-trash"></i> Remove
                    </button>
                </div>
            </div>
        `).join('');
    }
    
    updateSummary();
}

// Remove item from cart
function removeFromCart(itemId) {
    cart = cart.filter(item => item.id !== itemId);
    updateCart();
}

// Update order summary
function updateSummary() {
    const subtotal = cart.reduce((sum, item) => sum + item.price, 0);
    const tax = subtotal * 0.10;
    const total = subtotal + tax;
    
    document.getElementById('subtotal').textContent = `$${subtotal.toFixed(2)}`;
    document.getElementById('tax').textContent = `$${tax.toFixed(2)}`;
    document.getElementById('total').textContent = `$${total.toFixed(2)}`;
}

// Checkout
function checkout() {
    if (cart.length === 0) {
        alert('Your cart is empty!');
        return;
    }
    
    document.getElementById('order-modal').style.display = 'block';
}

// Close modal
function closeModal() {
    document.getElementById('order-modal').style.display = 'none';
}

// Set minimum date to today
function setMinDate() {
    const dateInput = document.getElementById('date');
    const today = new Date().toISOString().split('T')[0];
    dateInput.min = today;
}

// Submit order
function submitOrder(event) {
    event.preventDefault();
    
    const orderData = {
        customer: {
            name: document.getElementById('name').value,
            email: document.getElementById('email').value,
            phone: document.getElementById('phone').value
        },
        deliveryDate: document.getElementById('date').value,
        instructions: document.getElementById('message').value,
        items: cart,
        total: document.getElementById('total').textContent
    };
    
    // Here you would typically send the order to a server
    console.log('Order submitted:', orderData);
    
    alert(`Thank you for your order, ${orderData.customer.name}!\n\nOrder Total: ${orderData.total}\n\nWe will contact you soon to confirm your order.`);
    
    // Clear cart and close modal
    cart = [];
    updateCart();
    closeModal();
    document.getElementById('order-form').reset();
}

// Close modal when clicking outside
window.onclick = function(event) {
    const modal = document.getElementById('order-modal');
    if (event.target === modal) {
        closeModal();
    }
}