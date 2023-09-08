import axios from "axios";

export const useGetData = () => {

  const getData = async (info) => {
    let response = await axios.get(`http://localhost:8000/search?term=${info.value}`);
    console.log(response);
  };

  return { getData };
};