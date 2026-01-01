export namespace main {
	
	export class TaskOrder {
	    id: number;
	    order: number;
	
	    static createFrom(source: any = {}) {
	        return new TaskOrder(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.order = source["order"];
	    }
	}

}

export namespace models {
	
	export class Category {
	    id: number;
	    name: string;
	    exe_name: string;
	    is_active: boolean;
	
	    static createFrom(source: any = {}) {
	        return new Category(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.name = source["name"];
	        this.exe_name = source["exe_name"];
	        this.is_active = source["is_active"];
	    }
	}
	export class Task {
	    id: number;
	    name: string;
	    category_id: number;
	    is_completed: boolean;
	    // Go type: time
	    completed_at?: any;
	    reload_type: string;
	    reload_every: number;
	    reset_time: string;
	    reset_weekday?: number;
	    // Go type: time
	    last_done_at?: any;
	    // Go type: time
	    next_reset_at: any;
	    order: number;
	    // Go type: time
	    created_at: any;
	
	    static createFrom(source: any = {}) {
	        return new Task(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.name = source["name"];
	        this.category_id = source["category_id"];
	        this.is_completed = source["is_completed"];
	        this.completed_at = this.convertValues(source["completed_at"], null);
	        this.reload_type = source["reload_type"];
	        this.reload_every = source["reload_every"];
	        this.reset_time = source["reset_time"];
	        this.reset_weekday = source["reset_weekday"];
	        this.last_done_at = this.convertValues(source["last_done_at"], null);
	        this.next_reset_at = this.convertValues(source["next_reset_at"], null);
	        this.order = source["order"];
	        this.created_at = this.convertValues(source["created_at"], null);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}

}

