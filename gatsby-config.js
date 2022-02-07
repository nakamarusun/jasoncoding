module.exports = {
  siteMetadata: {
    title: `Jasoncoding`,
    siteUrl: `https://jasoncoding.com`,
  },
  plugins: [
    {
      resolve: "gatsby-plugin-react-svg",
      options: {
        rule: {
          include: /\.inline\.svg/,
          options: {
            classIdPrefix: true, // This does not work. Check out src/images/ltechs/scrubber.sh
          },
        },
      },
    },
    "gatsby-plugin-image",
    "gatsby-plugin-react-helmet",
    "gatsby-plugin-sitemap",
    "gatsby-plugin-sharp",
    "gatsby-transformer-sharp",
    "gatsby-plugin-postcss",
    "gatsby-plugin-fontawesome-css",
    {
      resolve: "gatsby-source-filesystem",
      options: {
        name: "images",
        path: "./src/images/",
      },
      __key: "images",
    },
    {
      resolve: `gatsby-plugin-manifest`,
      options: {
        name: "Jasoncoding",
        short_name: "Jasoncoding",
        start_url: "/",
        background_color: "#e5e7eb",
        theme_color: "#7e22ce",
        display: "standalone",
        icon: "src/images/icon.svg",
      },
    },
  ],
};
