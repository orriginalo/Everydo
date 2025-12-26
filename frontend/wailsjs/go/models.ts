export namespace models {
	
	export class Category {
	    id: number;
	    name: string;
	    exe_name: string;
	
	    static createFrom(source: any = {}) {
	        return new Category(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.name = source["name"];
	        this.exe_name = source["exe_name"];
	    }
	}

}

