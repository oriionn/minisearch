class SearchInput extends HTMLElement {
    constructor() {
        super();
        this.innerHTML = `
            <form action="/search" method="get" class="search">
                <input type="text" class="search-input" name="q" ${this.hasAttribute("value") ? `value="${this.getAttribute("value")}"`:""} />
                <input type="submit" class="search-input" value="🔍" />
            </form>
        `;
    }
}

customElements.define("search-input", SearchInput)
