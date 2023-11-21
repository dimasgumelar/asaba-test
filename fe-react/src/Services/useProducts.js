import { useEffect, useState } from "react";
import { axiosInstance } from "../Lib/Axios";

export default function useProducts() {
  const [products, setProducts] = useState([]);
  const [isLoading, setIsLoading] = useState(false);

  const getProducts = async () => {
    setIsLoading(true);
    try {
      const response = await axiosInstance.get("/barang");
      setProducts(response.data.data);
      setIsLoading(false);
    } catch (error) {
      console.log(error);
      setIsLoading(false);
    }
  };

  useEffect(() => {
    getProducts();
  }, []);
  return { data: products, isLoading };
}
