module.exports = {
    content: ["./src/**/*.{html,js}", "./html/**/*.html"],
    theme: {
        extend: {},
    },
    plugins: [
        require('@tailwindcss/forms'),
    ],
};