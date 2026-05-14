import pathlib
import re
import pymupdf

doc = pymupdf.open("raw/destination-cissp.pdf")
OUT_DIR = pathlib.Path("raw/destination-cissp")

# Ordered list of (pattern-to-match-first-line, section-slug).
# Domains 4/5/6 have no explicit "DOMAIN N" cover page — they're detected
# by the first numbered subsection (4.1 ..., 5.1 ..., 6.1 ...).
SECTION_RULES: list[tuple[re.Pattern, str]] = [
    (re.compile(r"^INTRODUCTION$"),  "introduction"),
    # TOC pages have "DOMAIN N: Title" (with colon); actual header pages have "DOMAIN N" alone.
    (re.compile(r"^DOMAIN 1$"),      "domain-1"),
    (re.compile(r"^DOMAIN 2$"),      "domain-2"),
    (re.compile(r"^DOMAIN 3$"),      "domain-3"),
    # Domains 4/5/6 have no standalone header page; detect by first subsection line.
    (re.compile(r"^4\.1\s"),         "domain-4"),
    (re.compile(r"^5\.1\s"),         "domain-5"),
    (re.compile(r"^6\.1\s"),         "domain-6"),
    (re.compile(r"^DOMAIN 7$"),      "domain-7"),
    (re.compile(r"^DOMAIN 8$"),      "domain-8"),
]


def detect_section(text: str, past_intro: bool) -> str | None:
    """Return the section slug if this page opens a new section, else None.

    `past_intro` gates detection of domain sections — TOC pages before the
    INTRODUCTION can also start with "4.1 ..." text, so we ignore them until
    we've confirmed the book's body has begun.
    """
    lines = [l.strip() for l in text.split("\n") if l.strip()]
    if not lines:
        return None
    first = lines[0]
    intro_pattern, intro_slug = SECTION_RULES[0]
    if intro_pattern.match(first):
        return intro_slug
    if not past_intro:
        return None
    for pattern, slug in SECTION_RULES[1:]:
        if pattern.match(first):
            return slug
    return None


def main() -> None:
    OUT_DIR.mkdir(parents=True, exist_ok=True)
    current_section: str | None = None
    past_intro = False

    for page_num, page in enumerate(doc, start=1):
        text = page.get_text()
        section = detect_section(text, past_intro)
        if section:
            current_section = section
            if section == "introduction":
                past_intro = True

        if current_section is None:
            continue  # front-matter before any recognised section

        section_dir = OUT_DIR / current_section
        section_dir.mkdir(parents=True, exist_ok=True)
        (section_dir / f"page-{page_num:04d}.md").write_text(text)
        print(f"[{current_section}] page {page_num}")


if __name__ == "__main__":
    main()
