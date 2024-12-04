<template>
    <div class="darModeButton">
        <label class="switchs-darkMode">
            <input type="checkbox" :checked="isDarkMode" @change="toggleDarkMode">
            <span class="sliders-darkMode"></span>
        </label>
    </div>
</template>

<script>
export default {
    data() {
        return { 
            isDarkMode: false,
        };
    },
    mounted() {
        // Récupère le statut du mode sombre dans localStorage
        const darkModeStatus = localStorage.getItem('isDarkMode');
        this.isDarkMode = darkModeStatus === 'true';

        // Applique le mode sombre en fonction du statut enregistré
        this.applyDarkMode(this.isDarkMode);
    },
    methods: {
        applyDarkMode(isDark) {
            if (isDark) {
                document.body.classList.add('darkmode');
            } else {
                document.body.classList.remove('darkmode');
            }

            // Applique la classe darkmode à tous les éléments si nécessaire
            const elements = document.querySelectorAll('*');
            elements.forEach(element => {
                if (isDark) {
                    element.classList.add('darkmode');
                } else {
                    element.classList.remove('darkmode');
                }
            });
        },
        toggleDarkMode() {
            this.isDarkMode = !this.isDarkMode;
            this.applyDarkMode(this.isDarkMode);

            // Sauvegarde l'état du mode sombre dans localStorage
            localStorage.setItem('isDarkMode', this.isDarkMode); 
        }
    }
};
</script>


