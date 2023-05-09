package constants

// Headers
const CONTENT_TYPE = "Content-Type"
const APPLICATION_JSON = "application/json"

// Domains
const EXITO_DOMAIN = "www.exito.com"
const EXITO_HALF_DOMAIN = "exito.com"
const AMAZON_DOMAIN = "www.amazon.com"
const AMAZON_HALF_DOMAIN = "amazon.com"
const JUMBO_DOMAIN = "www.tiendasjumbo.co"
const JUMBO_HALF_DOMAIN = "tiendasjumbo.co"

// Query Selectors
const EXITO_QUERY_SELECTOR = "div.flex.flex-grow-1.w-100.flex-column, div.flex.flex-column.min-vh-100.w-100"

// Amazon
const AMAZON_QUERY_SELECTOR = "div#centerCol.centerColAlign, li[data-csa-c-action='image-block-main-image-hover']"
const AMAZON_QUERY_NAME = "span[id='productTitle']"
const AMAZON_QUERY_BRAND = "div.celwidget  tbody tr.a-spacing-small.po-brand td.a-span9"
const AMAZON_QUERY_IMAGE_URL = "img.a-dynamic-image"
const AMAZON_QUERY_IMAGE_URL_ATTR = "data-old-hires"
const AMAZON_QUERY_DESCRIPTION = "div[id='feature-bullets'] ul li span.a-list-item"

const AMAZON_QUERY_CURRENTPRICE_DISCOUNT = "div.a-section.a-spacing-none.aok-align-center span.a-offscreen"
const AMAZON_QUERY_HIGHTPRICE_DISCOUNT = "div.a-section.a-spacing-small.aok-align-center span.a-offscreen"
const AMAZON_QUERY_DISCOUNT_DISCOUNT = "div.a-section.a-spacing-none.aok-align-center span.a-size-large.a-color-price.savingPriceOverride"

const AMAZON_QUERY_CURRENTPRICE_CURRENT = ".a-section.a-spacing-none.aok-align-center span.a-offscreen"
const AMAZON_QUERY_HIGHTPRICE_CURRENT = ".a-section.a-spacing-none.aok-align-center span.a-offscreen"

const AMAZON_QUERY_PRICES_TABLE = "div.a-section.a-spacing-small table.a-lineitem.a-align-top tr"
const AMAZON_QUERY_PRICE_ELEMENT_TABLE = "span.a-price.a-text-price.a-size-base span.a-offscreen ,span.a-price.a-text-price.a-size-medium.apexPriceToPay span.a-offscreen"

// Jumbo
const JUMBO_QUERY_SELECTOR = "div.vtex-store-components-3-x-productImage, div.vtex-flex-layout-0-x-flexCol.ml0.mr0.pl0.pr0.flex.flex-column.h-100.w-100"
const JUMBO_QUERY_NAME = "span.vtex-store-components-3-x-productBrand.vtex-store-components-3-x-productBrand--quickview"
const JUMBO_QUERY_IMAGE_URL = "img.vtex-store-components-3-x-productImageTag.vtex-store-components-3-x-productImageTag--main"
const JUMBO_QUERY_DESCRIPTION = "div.vtex-store-components-3-x-content.h-auto"
const JUMBO_QUERY_CURRENTPRICE = "div#items-price.flex.c-emphasis.tiendasjumboqaio-jumbo-minicart-2-x-cencoPrice div.tiendasjumboqaio-jumbo-minicart-2-x-price"
const JUMBO_QUERY_HIGHTPRICE = "div.b.ml2.tiendasjumboqaio-jumbo-minicart-2-x-cencoListPriceWrapper div.tiendasjumboqaio-jumbo-minicart-2-x-price"
const JUMBO_QUERY_DISCOUNT = "div.pr7.items-stretch.vtex-flex-layout-0-x-stretchChildrenWidth.flex span.vtex-product-price-1-x-currencyContainer.vtex-product-price-1-x-currencyContainer--summary"
const JUMBO_QUERY_PRODUCT_ID = "span.vtex-product-identifier-0-x-product-identifier__value"

// Cache
const CACHE = "./cache"
