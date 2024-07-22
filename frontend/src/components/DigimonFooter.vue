<template>
    <div class="footer">
        <div class="footer-content">
            <div class="social-links">
                <a :href="githubUrl" target="_blank" rel="noopener noreferrer">GitHub</a>
                <a :href="twitterUrl" target="_blank" rel="noopener noreferrer">Twitter</a>
                <a :href="linkedinUrl" target="_blank" rel="noopener noreferrer">LinkedIn</a>
            </div>
            <p>&copy; 2024 farizaotearoa</p>
        </div>
    </div>
</template>

<script>
export default {
    data() {
        return {
            githubUrl: 'https://github.com/farizaotearoa',
            twitterUrl: 'https://twitter.com/farizaotearoa',
            linkedinUrl: 'https://www.linkedin.com/in/farizaotearoarasyid/',
        };
    },
};
</script>

<style scoped>
.footer {
    background-color: #fff;
    color: #000;
    font-weight: bold;
    padding: 1rem;
    text-align: center;
}

.footer-content {
    max-width: 1200px;
    margin: 0 auto;
}

.social-links a {
    color: #000;
    margin: 0 10px;
    font-weight: bold;
    text-decoration: none;
}

.social-links a:hover {
    text-decoration: underline;
}
</style>