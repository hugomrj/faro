export namespace categories {
	
	export class Category {
	    id: number;
	    name: string;
	    description: string;
	
	    static createFrom(source: any = {}) {
	        return new Category(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.name = source["name"];
	        this.description = source["description"];
	    }
	}
	export class CategoryCreateRequest {
	    name: string;
	    description: string;
	
	    static createFrom(source: any = {}) {
	        return new CategoryCreateRequest(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.name = source["name"];
	        this.description = source["description"];
	    }
	}
	export class CategoryUpdateRequest {
	    name: string;
	    description: string;
	
	    static createFrom(source: any = {}) {
	        return new CategoryUpdateRequest(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.name = source["name"];
	        this.description = source["description"];
	    }
	}

}

export namespace products {
	
	export class CreateRequest {
	    name: string;
	    category_id: number;
	    price: number;
	    stock: number;
	    status: string;
	
	    static createFrom(source: any = {}) {
	        return new CreateRequest(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.name = source["name"];
	        this.category_id = source["category_id"];
	        this.price = source["price"];
	        this.stock = source["stock"];
	        this.status = source["status"];
	    }
	}
	export class Product {
	    id: number;
	    name: string;
	    category_id: number;
	    price: number;
	    stock: number;
	    status: string;
	    created_at: string;
	    category_name?: string;
	
	    static createFrom(source: any = {}) {
	        return new Product(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.name = source["name"];
	        this.category_id = source["category_id"];
	        this.price = source["price"];
	        this.stock = source["stock"];
	        this.status = source["status"];
	        this.created_at = source["created_at"];
	        this.category_name = source["category_name"];
	    }
	}

}

