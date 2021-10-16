package product

//
//func (a Product) CreateProductDetail(c *gin.Context) {
//	metadata, errA := a.AuthUsecase.ExtractTokenMetadata(c.Request)
//	if errA != nil {
//		c.JSON(http.StatusBadRequest, gin.H{
//			"message": errA.Error(),
//		})
//		return
//	}
//	fmt.Println(metadata)
//	isAdmin := a.AccountUsecase.CheckUserIsAdmin(metadata.Email)
//	if isAdmin {
//		product_id := c.PostForm("product_id")
//		size_id := c.PostForm("size_id")
//		price := c.PostForm("price")
//		description := c.PostForm("description")
//
//		if product_id == "" || size_id == "" || price == "" || description = "" {
//			c.JSON(http.StatusBadRequest, gin.H{
//				"message": "product_id, size_id, price, and description(255 chars) cannot be Empty",
//			})
//			return
//		}
//		productid, errproductid := strconv.Atoi(product_id)
//		sizeid, errsizeid := strconv.Atoi(size_id)
//		priceint, errpriceint := strconv.Atoi(price)
//
//		form_product := models.ProductDetail{
//			ProductId: ,
//			ProductId: ,
//			Description: description,
//			Price: price,
//		}
//		response := a.ProductUsecase.CreateProductDetail(form_product)
//		c.JSON(response.Status, response)
//		return
//	}
//	c.JSON(http.StatusBadRequest, gin.H{
//		"message": "you are not allowed",
//	})
//
//}

// func (a Product) UpdateProductDetail(c *gin.Context) {
// 	metadata, errA := a.AuthUsecase.ExtractTokenMetadata(c.Request)
// 	if errA != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"message": errA.Error(),
// 		})
// 		return
// 	}
// 	fmt.Println(metadata)
// 	isAdmin := a.AccountUsecase.CheckUserIsAdmin(metadata.Email)
// 	if isAdmin {
// 		name := c.PostForm("name")
// 		category_id := c.PostForm("category_id")

// 		if name == "" || category_id == "" {
// 			c.JSON(http.StatusBadRequest, gin.H{
// 				"message": "name and category_id cannot be Empty",
// 			})
// 			return
// 		}

// 		tint, _ := strconv.Atoi(category_id)

// 		form_product := models.ProductCategory{
// 			Name: name,
// 			Id:   tint,
// 		}
// 		response := a.ProductUsecase.UpdateProductDetail(form_product)
// 		c.JSON(response.Status, response)
// 		return
// 	}
// 	c.JSON(http.StatusBadRequest, gin.H{
// 		"message": "you are not allowed",
// 	})
// }
