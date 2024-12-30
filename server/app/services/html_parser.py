from bs4 import BeautifulSoup


def parse_html(content: bytes):
    html_content = content.decode("utf-8")
    soup = BeautifulSoup(html_content, "html.parser")
    headers = [header.get_text() for header in soup.find_all("h1")]
    return headers