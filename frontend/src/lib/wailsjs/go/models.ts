export namespace main {
	
	export class Maker {
	    id: string;
	    name: string;
	
	    static createFrom(source: any = {}) {
	        return new Maker(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.name = source["name"];
	    }
	}
	export class Model {
	    id: string;
	    makerId: string;
	    name: string;
	    year: number;
	
	    static createFrom(source: any = {}) {
	        return new Model(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.makerId = source["makerId"];
	        this.name = source["name"];
	        this.year = source["year"];
	    }
	}
	export class Vehicle {
	    id: string;
	    modelId: string;
	    vehicleType: string;
	    plateNo: string;
	    fuelType: string;
	    capacity: number;
	    deptAssigned: string;
	
	    static createFrom(source: any = {}) {
	        return new Vehicle(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.modelId = source["modelId"];
	        this.vehicleType = source["vehicleType"];
	        this.plateNo = source["plateNo"];
	        this.fuelType = source["fuelType"];
	        this.capacity = source["capacity"];
	        this.deptAssigned = source["deptAssigned"];
	    }
	}

}

