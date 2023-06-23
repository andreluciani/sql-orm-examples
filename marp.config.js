const canonicalUrl = process.env.URL || undefined
const ogImage = (() => {
    if (canonicalUrl) return `${canonicalUrl}/og-image.jpg`
    return undefined
})()

module.exports = {
    allowLocalFiles: true,
    ogImage,
    themeSet: 'themes',
    url: canonicalUrl,
}