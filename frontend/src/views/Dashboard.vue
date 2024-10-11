<template>
    <div>
        <h2>Welcome, {{ user.fullname }}!</h2>
        <p>Username: {{ user.username }}</p>
        <button @click="logout">Logout</button>
    </div>
</template>

<script>
import axios from "axios";

export default {
    data() {
        return {
            user: JSON.parse(localStorage.getItem("user") || "{}"),
        };
    },
    methods: {
        async logout() {
            try {
                await axios.post("/logout");
                localStorage.removeItem("user");
                this.$router.push("/login");
            } catch (error) {
                console.error("Logout failed", error);
            }
        },
    },
};
</script>
