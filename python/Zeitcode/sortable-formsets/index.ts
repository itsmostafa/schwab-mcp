import Mustache from "mustache";

// Select the parent ul element
const parentElement = document.querySelector<HTMLElement>("#GroupBy");

// initialize list element count
let liCount = 0;

if (parentElement) {
    // update element count with length of ul elements
    liCount = parentElement.children.length;

    for (const currElem of parentElement.querySelectorAll("li")) {
        const upElem = currElem.querySelector(".up");
        const downElem = currElem.querySelector(".down");
        const removeElem = currElem.querySelector(".remove");

        // add click listener to up buttons
        if (upElem) {
            upElem.addEventListener("click", event =>
                moveUp(event, currElem, parentElement)
            );
        }

        // add click listener to down buttons
        if (downElem) {
            downElem.addEventListener("click", event =>
                moveDown(event, currElem, parentElement)
            );
        }

        // add click listener to remove button
        if (removeElem) {
            removeElem.addEventListener("click", event =>
                removeFromList(event, currElem, parentElement)
            );
        }
    }

    const addButton = document.querySelector("#add");

    // add a new list element by appending it to the end of the list
    if (addButton) {
        addButton.addEventListener("click", event => {
            const templateElem = document.querySelector("#litemplate");

            if (templateElem) {
                const liContent = Mustache.render(templateElem.innerHTML, {
                    number: ++liCount
                });

                const newerElem = document.createElement("li");

                newerElem.innerHTML = liContent;

                const newerUpElem = newerElem.querySelector(".up");

                if (newerUpElem) {
                    newerUpElem.addEventListener("click", event =>
                        moveUp(event, newerElem, parentElement)
                    );
                }

                const newerDownElem = newerElem.querySelector(".down");

                if (newerDownElem) {
                    newerDownElem.addEventListener("click", event =>
                        moveDown(event, newerElem, parentElement)
                    );
                }

                const newerRemoveElem = newerElem.querySelector(".remove");

                if (newerRemoveElem) {
                    newerRemoveElem.addEventListener("click", event =>
                        removeFromList(event, newerElem, parentElement)
                    );
                }

                parentElement.insertBefore(newerElem, parentElement.firstChild);
                recalculateWeight(parentElement);
            }
        });
    }
}

// move list element up one
function moveUp(event: Event, liElem: HTMLElement, parentElem: HTMLElement) {
    const prevElem = liElem.previousElementSibling;
    if (prevElem) {
        parentElem.insertBefore(liElem, prevElem);
        recalculateWeight(parentElem);
    }
}

// move list element down one
function moveDown(event: Event, liElem: HTMLElement, parentElem: HTMLElement) {
    const nextElem = liElem.nextElementSibling;
    if (nextElem) {
        parentElem.insertBefore(nextElem, liElem);
        recalculateWeight(parentElem);
    }
}

// remove an element from the list
function removeFromList(
    event: Event,
    liElem: HTMLElement,
    parentElem: HTMLElement
) {
    if (liElem) {
        parentElem.removeChild(liElem);
        liCount--;
        recalculateWeight(parentElem);
    }
}

// recalculate list element indices
function recalculateWeight(parentElem: HTMLElement) {
    for (let i = 0; i < parentElem.children.length; i++) {
        const weightElem = parentElem.children[i].querySelector<
            HTMLInputElement
        >("[name='weight']");
        if (weightElem) {
            weightElem.value = i.toString();
        }
    }
}
