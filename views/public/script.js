async function fetchServices() {
    try {
        const response = await fetch("/api");
        const data = await response.json();

        const list = document.getElementById("service-list");
        list.innerHTML = "";

        data.forEach(service => {
         const card = `
    <div class="card">
        <div class="info">
            <span class="dot ${service.status}"></span>
            <span class="service">${service.service}</span>
        </div>

        <div class="badge ${service.status}">
            ${service.status}
        </div>
    </div>
`;
            list.innerHTML += card;
        });

    } catch (error) {
        console.error("Erro ao buscar serviços:", error);
    }
}


setInterval(fetchServices, 10000);