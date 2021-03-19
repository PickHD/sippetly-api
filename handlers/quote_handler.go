package handlers

import (
	"context"
	"net/http"
	"time"
	"strconv"
	"sippetly-api/db"
	"sippetly-api/models"
	

	"github.com/gin-gonic/gin"
)

func CreateQuoteHandler(c *gin.Context){
	var newQuote models.Quote
	
	if err:=c.ShouldBindJSON(&newQuote);err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{
			"success":false,
			"err_code":http.StatusBadRequest,
			"err_message":err.Error(),
		})
		return
	}

	db.OpenDB()

	query:="INSERT INTO quotes(quote_name,created_by) VALUES (?,?)"

	ctx,cancelFunc:=context.WithTimeout(context.Background(),5*time.Second)
	defer cancelFunc()

	stmt,err:=db.DB.PrepareContext(ctx,query)
	if err!=nil{
		c.JSON(http.StatusInternalServerError,gin.H{
			"success":false,
			"err_code":http.StatusInternalServerError,
			"err_message":err.Error(),
		})
		return
	}
	defer stmt.Close()

	res,err:=stmt.ExecContext(ctx,newQuote.QuoteName,newQuote.CreatedBy)
	if err!=nil{
		c.JSON(http.StatusInternalServerError,gin.H{
			"success":false,
			"err_code":http.StatusInternalServerError,
			"err_message":err.Error(),
		})
		return
	}

	rows,err:=res.RowsAffected()
	if err!=nil{
		c.JSON(http.StatusInternalServerError,gin.H{
			"success":false,
			"err_code":http.StatusInternalServerError,
			"err_message":err.Error(),
		})
		return
	}

	db.CloseDB()

	c.JSON(http.StatusCreated,gin.H{
		"success":true,
		"code":http.StatusCreated,
		"message":"Quotes Created Successfully",
		"result":rows,
	})

}
func RetrieveAllQuoteHandler(c *gin.Context){
	var getAllQuotes []*models.Quote

	db.OpenDB()

	ctx,cancelFunc:=context.WithTimeout(context.Background(),5*time.Second)
	defer cancelFunc()

	rows,err:=db.DB.QueryContext(ctx,"SELECT * FROM quotes;")
	if err!=nil{
		c.JSON(http.StatusInternalServerError,gin.H{
			"success":false,
			"err_code":http.StatusInternalServerError,
			"err_message":err.Error(),
		})
		return
	}
	defer rows.Close()

	for rows.Next(){
		q:=new(models.Quote)
		err := rows.Scan(&q.Id,&q.QuoteName,&q.CreatedBy,&q.CreatedAt,&q.UpdatedAt)

		if err != nil {
			c.JSON(http.StatusInternalServerError,gin.H{
				"success":false,
				"err_code":http.StatusInternalServerError,
				"err_message":err.Error(),
			})
			return
		}

		getAllQuotes=append(getAllQuotes, q)

	}
	err = rows.Err()
	if err!=nil{
		c.JSON(http.StatusInternalServerError,gin.H{
			"success":false,
			"err_code":http.StatusInternalServerError,
			"err_message":err.Error(),
		})
		return
	}
	db.CloseDB()

	c.JSON(http.StatusOK,gin.H{
			"success":true,
			"code":http.StatusOK,
			"message":nil,
			"result":	getAllQuotes,
	})

}
func GetOneQuoteHandler(c *gin.Context){
	var oneQuote models.Quote
	
	quoteId,err:=strconv.Atoi(c.Param("id"))
	if err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{
			"success":false,
			"err_code":http.StatusBadRequest,
			"err_message":err.Error(),
		})
		return
	}

	db.OpenDB()

	err=db.DB.QueryRow("SELECT * FROM quotes WHERE id=?",quoteId).Scan(&oneQuote.Id,&oneQuote.QuoteName,&oneQuote.CreatedBy,&oneQuote.CreatedAt,&oneQuote.UpdatedAt)

	if err!=nil{
		c.JSON(http.StatusNotFound,gin.H{
			"success":false,
			"err_code":http.StatusNotFound,
			"err_message":"Quotes not found",
		})
		return
	}

	db.CloseDB()

	c.JSON(http.StatusOK,gin.H{
		"success":true,
		"code":http.StatusOK,
		"message":nil,
		"result":oneQuote,
	})

}
func UpdateQuoteHandler(c *gin.Context){
	var updateQuote models.Quote
	
	quoteId,err:=strconv.Atoi(c.Param("id"))
	if err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{
			"success":false,
			"err_code":http.StatusBadRequest,
			"err_message":err.Error(),
		})
		return
	}

	if err:=c.BindJSON(&updateQuote);err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{
			"success":false,
			"err_code":http.StatusBadRequest,
			"err_message":err.Error(),
		})
		return
	}

	db.OpenDB()

	ctx,cancelFunc:=context.WithTimeout(context.Background(),5*time.Second)
	defer cancelFunc()

	query:="UPDATE quotes SET quote_name=?,created_by=? WHERE id=?"

	stmt,err:=db.DB.PrepareContext(ctx,query)
	if err!=nil{
		c.JSON(http.StatusInternalServerError,gin.H{
			"success":false,
			"err_code":http.StatusInternalServerError,
			"err_message":err.Error(),
		})
		return
	}
	defer stmt.Close()

	res,err:=stmt.ExecContext(ctx,updateQuote.QuoteName,updateQuote.CreatedBy,quoteId)
	if err!=nil{
		c.JSON(http.StatusInternalServerError,gin.H{
			"success":false,
			"err_code":http.StatusInternalServerError,
			"err_message":err.Error(),
		})
		return
	}
	if rows,_:=res.RowsAffected();rows==0{
		c.JSON(http.StatusNotFound,gin.H{
			"success":false,
			"err_code":http.StatusNotFound,
			"err_message":"Qoutes not found.",
		})
		return
	}

	db.CloseDB()

	c.JSON(http.StatusNonAuthoritativeInfo,gin.H{
		"success":true,
		"code":http.StatusNonAuthoritativeInfo,
		"message":"Quotes Updated Successfully",
		"result":nil,
	})
	
}
func DeleteQuoteHandler(c *gin.Context){	
	quoteId,err:=strconv.Atoi(c.Param("id"))
	if err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{
			"success":false,
			"err_code":http.StatusBadRequest,
			"err_message":err.Error(),
		})
		return
	}

	db.OpenDB()
	
	ctx,cancelFunc:=context.WithTimeout(context.Background(),5*time.Second)
	defer cancelFunc()

	query:="DELETE FROM quotes WHERE id=?"

	stmt,err:=db.DB.PrepareContext(ctx,query)
	if err!=nil{
		c.JSON(http.StatusInternalServerError,gin.H{
			"success":false,
			"err_code":http.StatusInternalServerError,
			"err_message":err.Error(),
		})
		return
	}
	defer stmt.Close()

	rows,err:=stmt.ExecContext(ctx,quoteId)
	if res,_:=rows.RowsAffected();res==0{
		c.JSON(http.StatusNotFound,gin.H{
			"success":false,
			"err_code":http.StatusNotFound,
			"err_message":"Qoutes not found.",
		})
		return
	}
	if err!=nil{
		c.JSON(http.StatusInternalServerError,gin.H{
			"success":false,
			"err_code":http.StatusInternalServerError,
			"err_message":err.Error(),
		})
		return
	}
	db.CloseDB()

	c.JSON(http.StatusOK,gin.H{
		"success":true,
		"code":http.StatusOK,
		"message":"Quotes Deleted Successfully",
		"result":nil,
	})
}