from bs4 import BeautifulSoup


def parse_html(content: bytes):
    """
    Parses HTML content and extracts all h1 headers.

    This function takes in HTML content as bytes, decodes it to a UTF-8 string,
    and uses BeautifulSoup to parse the HTML. It extracts text from all h1 
    header elements and returns them as a list of strings.

    :param content: HTML content in byte format.
    :return: A list of strings containing the text of all h1 headers.
    """

    html_content = content.decode("utf-8")
    soup = BeautifulSoup(html_content, "html.parser")
    headers = [header.get_text() for header in soup.find_all("h1")]
    return headers
