export const fetchData = async (api, params) => {
  const queryParam = getQueryParam(params)
  const url = import.meta.env.VITE_API_URL + api + '?' + queryParam
  const res = await fetch(url)
  if (!res.ok) {
    throw new Error('Network response was not ok');
  }
  return res.json();
}

export const getQueryParam = (params) => {
  return new URLSearchParams(params)
}