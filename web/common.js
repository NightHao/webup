const api = window.location.protocol + "//" + window.location.host + ":6101";

let lang;
let toggleLegacyIsShort = false;

function addParamToUrl(url, key, val) {
    const u = new URL(url);
    u.searchParams.set(key, val);
    return u.toString();
}

function addMenuItem(parent, item) {
    const link = document.createElement("a");
    link.href = item.link;
    link.innerText = item.title;
    link.className = "main-menu";
    link.href = addParamToUrl(link.href, "lang", lang);
    if (item.type === "DOCLIST") {
        link.href = addParamToUrl(link.href, "c_id", item.data);
    }
    const li = document.createElement("li");
    li.appendChild(link);
    parent.appendChild(li);
    return { // returning the element for legacy customisation TODO: refactor this..?
        li: li,
        a: link,
    };
}

function addMenuItemWrapper(parent, item)
{
    const r = addMenuItem(parent, item);
    if (item.children !== null) {
        const dropdown = r;
        dropdown.li.className = "dropdown";
        dropdown.a.className = "main-menu";
        dropdown.a.href = "javascript:void(0)";
        const arrow = document.createElement("span");
        arrow.className = "fa fa-angle-down icons-dropdown";
        dropdown.a.appendChild(arrow);
        const sub = document.createElement("ul");
        sub.className = "dropdown-menu edugate-dropdown-menu-1";
        if (toggleLegacyIsShort) {
            sub.style.top = "55px";
        }
        item.children.forEach(subItem => {
            const elem = addMenuItem(sub, subItem);
            elem.li.className = "";
            elem.a.className = "link-page";
        });
        dropdown.li.appendChild(sub);
    }
    return r;
}

// TODO: when fetch failed, add legacy menu ...?

async function load() {
    const menu = await (await fetch(api + "/cms/menu/" + lang)).json();
    const menuUI = document.getElementById("menu");
    // addLegacyMenu(menuUI);
    menu.forEach(item => addMenuItemWrapper(menuUI, item));

    document.getElementById("logoText").innerText = lang === "en" ?
        "Artificial Intelligence Talent Cultivation Project" :
        "教育部人工智慧技術及應用人才培育計畫";
}
