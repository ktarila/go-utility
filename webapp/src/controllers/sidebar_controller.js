import { Controller } from "@hotwired/stimulus";
import { useClickOutside } from "stimulus-use";
import PerfectScrollbar from "perfect-scrollbar";

export default class extends Controller {
    static get targets() {
        return ["draweraside", "drawertransition", "navtree"];
    }

    connect() {
        useClickOutside(this);
        this.drawer();
        const ps = new PerfectScrollbar("#scroll-container");
    }

    initialize() {
        this.open = true;
    }

    clickOutside(event) {
        // example to close a modal
        // this.toggleableTarget.classList.add('hidden')
        // this.open = false;
        // this.drawer();
    }

    closeSidebar() {
        event.preventDefault();
        this.open = false;
        this.drawer();
    }

    toggleSidebar()
    {
        event.preventDefault();
        if (this.open == true)
        {
            this.closeSidebar();
        } else {
            this.openSidebar();
        }
    }

    resetOpen() {
        this.closeSidebar();
    }

    openSidebar() {
        this.open = true;
        this.drawer();
    }

    drawer() {
        if (this.open) {
            document.body.style.setProperty("overflow", "hidden");
            this.drawertransitionTarget.classList.remove("hidden");
            this.drawerasideTarget.classList.remove("absolute");
            this.drawerasideTarget.classList.add("translate-x-0");
            this.drawerasideTarget.classList.remove("-translate-x-full");
        } else {
            document.body.style.removeProperty("overflow");
            this.drawertransitionTarget.classList.add("hidden");
            this.drawerasideTarget.classList.add("absolute");
            this.drawerasideTarget.classList.remove("translate-x-0");
            this.drawerasideTarget.classList.add("-translate-x-full");
        }
    }

    showTree(event) {
        const tree_root = event.currentTarget;
        const tree = tree_root.children[1];
        tree.classList.toggle("hidden");
        const svgContainer = tree_root.children[0].children[1];
        svgContainer.children[0].classList.toggle("hidden");
        svgContainer.children[1].classList.toggle("hidden");
    }
}
