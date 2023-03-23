const api = window.location.protocol + "//" + window.location.host + ":6101";

let lang;

function addMenuItem(parent, item) {
    const link = document.createElement("a");
    link.href = item.link;
    link.innerText = item.label;
    link.className = "main-menu";
    const li = document.createElement("li");
    li.appendChild(link);
    parent.appendChild(li);
    return { // returning the element for legacy customisation
        li: li,
        a: link,
    };
}

function addLegacyMenu(parent) {
    const items = lang === "en" ?
        [
            {
                label: "HOME",
                link: "index.html",
            },
            {
                label: "COURSES",
                link: "courses.html",
            },
            {
                label: "ROADMAP",
                dropdown: [
                    {
                        label: "Static Roadmap",
                        link: "roadmap-static.html",
                    },
                    {
                        label: "Interactive Roadmap",
                        link: "roadmap-interactive.html",
                    },
                ]
            },
        ] :
        [
            {
                label: "首頁",
                link: "index-chinese.html",
            },
            {
                label: "課程",
                link: "courses-chinese.html",
            },
            {
                label: "課程地圖",
                dropdown: [
                    {
                        label: "靜態式地圖",
                        link: "roadmap-static-chinese.html",
                    },
                    {
                        label: "互動式地圖",
                        link: "roadmap-interactive-chinese.html",
                    },
                ]
            },
        ];

    addMenuItem(parent, items[0]);
    addMenuItem(parent, items[1]);
    const dropdown = addMenuItem(parent, items[2]);
    dropdown.li.className = "dropdown";
    dropdown.a.className = "main-menu";
    dropdown.a.href = "javascript:void(0)";
    const arrow = document.createElement("span");
    arrow.className = "fa fa-angle-down icons-dropdown";
    dropdown.a.appendChild(arrow);
    const sub = document.createElement("ul");
    sub.className = "dropdown-menu edugate-dropdown-menu-1";
    items[2].dropdown.forEach(subItem => {
        const elem = addMenuItem(sub, subItem);
        elem.li.className = "";
        elem.a.className = "link-page";
    });
    dropdown.li.appendChild(sub);
}

async function load() {
    const menu = await (await fetch(api + "/cms/menu/" + lang)).json();
    const menuUI = document.getElementById("menu");
    addLegacyMenu(menuUI);
    menu.forEach(item => addMenuItem(menuUI, item));

    document.getElementById("logoText").innerText = lang === "en" ?
        "Artificial Intelligence Talent Cultivation Project" :
        "教育部人工智慧技術及應用人才培育計畫";
}
